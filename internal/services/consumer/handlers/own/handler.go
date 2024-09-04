package own

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/mailru/easyjson"
	"github.com/oschwald/geoip2-golang"
	kafkaLib "github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/pkg/domain/event"
)

type EventRepository interface {
	AddEvent(
		ctx context.Context,
		event *event.EventExtended,
	) error
}

type Metrics interface {
	IncConsumerError(err error)
}

type Handler struct {
	city            *geoip2.Reader
	asn             *geoip2.Reader
	eventRepository EventRepository

	metrics Metrics

	withGeo bool
}

func (h *Handler) Handle(ctx context.Context, message kafkaLib.Message) error {
	now := time.Now()

	evt := &event.EventExtended{}
	if err := easyjson.Unmarshal(message.Value, evt); err != nil {
		h.metrics.IncConsumerError(err)

		return fmt.Errorf("consumer proc: unmarshal json %w", err)
	}

	evt.FromQueueDatetime = now

	if h.withGeo && evt.IP != "" {
		nIP := net.ParseIP(evt.IP)
		if nIP != nil {
			h.enrichEventLocation(evt, nIP)
		}
	}

	if err := h.eventRepository.AddEvent(ctx, evt); err != nil {
		return fmt.Errorf("consumer proc: add event: %w", err)
	}

	return nil
}

func (h *Handler) enrichEventLocation(record *event.EventExtended, nIP net.IP) {
	var (
		err  error
		city *geoip2.City
		as   *geoip2.ASN
	)

	if city, err = h.city.City(nIP); err == nil {
		record.CountryIsoCode = city.Country.IsoCode

		cityFound := false
		if name, ok := city.City.Names["en"]; ok {
			record.City = name
			cityFound = true
		}

		if !cityFound {
			if name, ok := city.City.Names["EN"]; ok {
				record.City = name
				cityFound = true
			}
		}

		if !cityFound && len(city.City.Names) > 0 {
			// peek first city name, unordered
			for k := range city.City.Names {
				record.City = city.City.Names[k]
				break
			}
		}

		if len(city.Subdivisions) > 0 {
			record.Region = city.Subdivisions[0].IsoCode
		}
	}

	if as, err = h.asn.ASN(nIP); err == nil {
		record.AS = uint32(as.AutonomousSystemNumber)
		record.ORG = as.AutonomousSystemOrganization
	}
}

func New(city, asn *geoip2.Reader, eventRepository EventRepository, metrics Metrics, withGeo bool) *Handler {
	return &Handler{
		city:            city,
		asn:             asn,
		eventRepository: eventRepository,
		metrics:         metrics,
		withGeo:         withGeo,
	}
}
