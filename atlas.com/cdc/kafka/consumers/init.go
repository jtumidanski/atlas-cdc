package consumers

import (
	"atlas-cdc/kafka/handler"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	DamageCharacterCommand = "damage_character_command"
)

func CreateEventConsumers(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, DamageCharacterCommand, topicToken, emptyEventCreator, processor)
	}
	cec("DAMAGE_CHARACTER", DamageCharacterCommandCreator(), HandleDamageCharacterCommand())
}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, name string, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, name, topicToken, "Character Damage Coordinator", emptyEventCreator, processor)
}
