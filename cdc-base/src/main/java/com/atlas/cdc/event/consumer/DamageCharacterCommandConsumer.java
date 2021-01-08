package com.atlas.cdc.event.consumer;

import com.atlas.cdc.command.DamageCharacterCommand;
import com.atlas.cdc.constant.CommandConstants;
import com.atlas.cdc.processor.CharacterDamageProcessor;
import com.atlas.cdc.processor.TopicDiscoveryProcessor;
import com.atlas.kafka.consumer.SimpleEventHandler;

public class DamageCharacterCommandConsumer implements SimpleEventHandler<DamageCharacterCommand> {
   @Override
   public void handle(Long key, DamageCharacterCommand command) {
      CharacterDamageProcessor.process(command.characterId(), command.monsterId(), command.monsterUniqueId(), command.damage(),
            command.damageFrom(), command.element(), command.direction());
   }

   @Override
   public Class<DamageCharacterCommand> getEventClass() {
      return DamageCharacterCommand.class;
   }

   @Override
   public String getConsumerId() {
      return "Damage Character Coordinator";
   }

   @Override
   public String getBootstrapServers() {
      return System.getenv("BOOTSTRAP_SERVERS");
   }

   @Override
   public String getTopic() {
      return TopicDiscoveryProcessor.getTopic(CommandConstants.DAMAGE_CHARACTER);
   }
}
