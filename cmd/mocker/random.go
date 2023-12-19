package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type randomMachine struct {
	countries          []string
	cities             []string
	appsIDs            []int
	installationIDs    []string
	sessionIDs         []string
	appVersions        []string
	deviceModels       []string
	deviceTypes        []string
	deviceLocales      []string
	deviceManufactures []string
	osNames            []string
	osVersions         []string
	eventNames         []string
	packages           []string
	appMetricaIDs      []string
}

func rnd() *rand.Rand {
	s := rand.NewSource(time.Now().Unix())
	return rand.New(s)
}

func (r *randomMachine) rndAppID() int {
	return r.appsIDs[rnd().Intn(len(r.appsIDs))]
}

func (r *randomMachine) rndOsName() string {
	return r.osNames[rnd().Intn(len(r.osNames))]
}

func (r *randomMachine) rndOsVersion() string {
	return r.osVersions[rnd().Intn(len(r.osVersions))]
}

func (r *randomMachine) rndDeviceManufacturer() string {
	return r.deviceManufactures[rnd().Intn(len(r.deviceManufactures))]
}

func (r *randomMachine) rndDeviceModel() string {
	return r.deviceModels[rnd().Intn(len(r.deviceModels))]
}

func (r *randomMachine) rndCountry() string {
	return r.countries[rnd().Intn(len(r.countries))]
}

func (r *randomMachine) rndCity() string {
	return r.cities[rnd().Intn(len(r.cities))]
}

func (r *randomMachine) rndInstallationID() string {
	return r.installationIDs[rnd().Intn(len(r.installationIDs))]
}

func (r *randomMachine) rndSessionID() string {
	return r.sessionIDs[rnd().Intn(len(r.sessionIDs))]
}

func (r *randomMachine) rndAppVersion() string {
	return r.appVersions[rnd().Intn(len(r.appVersions))]
}

func (r *randomMachine) rndDeviceType() string {
	return r.deviceTypes[rnd().Intn(len(r.deviceTypes))]
}

func (r *randomMachine) rndDeviceLocale() string {
	return r.deviceLocales[rnd().Intn(len(r.deviceLocales))]
}

func (r *randomMachine) rndEventName() string {
	return r.eventNames[rnd().Intn(len(r.eventNames))]
}

func (r *randomMachine) rndField() string {
	return "rndFieldName()"
}

func (r *randomMachine) rndMethod() string {
	return "rndMethodName()"
}

func (r *randomMachine) rndPackage() string {
	return r.packages[rnd().Intn(len(r.packages))]
}

func (r *randomMachine) rndAppmetricaDeviceID() string {
	return r.appMetricaIDs[rnd().Intn(len(r.appMetricaIDs))]
}

func (r *randomMachine) rndIosIfa() string {
	return generateRandomUUID()
}

func (r *randomMachine) rndIosIfv() string {
	return generateRandomUUID()
}

func (r *randomMachine) rndAndroidID() string {
	return generateRandomUUID()
}

func (r *randomMachine) rndGoogleAid() string {
	return generateRandomUUID()
}

func (r *randomMachine) rndProfileID() string {
	return generateRandomUUID()
}

func (r *randomMachine) rndEventJSON() string {
	return generateRandomJSON()
}

func (r *randomMachine) rndEventDatetime() string {
	return generateRandomDateTime()
}

func (r *randomMachine) rndEventTimestamp() int {
	return int(generateRandomTimestamp())
}

func (r *randomMachine) rndEventReceiveDatetime() string {
	return generateRandomDateTime()
}

func (r *randomMachine) rndEventReceiveTimestamp() string {
	return fmt.Sprintf("%d", generateRandomTimestamp())
}

func (r *randomMachine) rndConnectionType() string {
	return r.deviceTypes[rnd().Intn(len(r.deviceTypes))]
}

func (r *randomMachine) rndOperatorName() string {
	return generateRandomString(10)
}

func (r *randomMachine) rndMcc() string {
	return generateRandomString(3)
}

func (r *randomMachine) rndMnc() string {
	return generateRandomString(3)
}

func (r *randomMachine) rndAppPackageName() string {
	return generateRandomString(10)
}

func generateRandomUUID() string {
	uuidObj, _ := uuid.NewRandom()
	return uuidObj.String()
}

func generateRandomJSON() string {
	return fmt.Sprintf(`{
   "summary":[
      {
         "платформа": "android.tv"
      },
      {
         "программа": "%s %s"
      },
      {
         "idканала": "%d"
      },
      {
         "источник": "%s"
      }
   ],
   "%s:%d": "онлайн"
}`,
		gofakeit.DateRange(
			time.Now().Add((24*time.Hour)*-40),
			time.Now(),
		).Format("02.01.06, 15:04"),
		gofakeit.MovieName(),
		gofakeit.Int16(),
		gofakeit.MovieGenre(),
		gofakeit.LetterN(10),
		gofakeit.Int16(),
	)
}

func generateRandomDateTime() string {
	return gofakeit.DateRange(
		time.Now().Add((24*time.Hour)*-40),
		time.Now(),
	).Format("2006-01-02 15:04:05")
}

func generateRandomTimestamp() int64 {
	return gofakeit.DateRange(
		time.Now().Add((24*time.Hour)*-40),
		time.Now(),
	).Unix()
}

func generateRandomString(length uint) string {
	return gofakeit.LetterN(length)
}

func generateStrings(amount int, generator func() string, length ...uint) []string {
	var strings []string
	for i := 0; i < amount; i++ {
		if len(length) > 0 {
			strings = append(strings, gofakeit.LetterN(length[0]))
		} else {
			strings = append(strings, generator())
		}
	}
	return strings
}

func generateIntegers(amount, min, max int) []int {
	var integers []int
	for i := 0; i < amount; i++ {
		integers = append(integers, gofakeit.Number(min, max))
	}
	return integers
}
