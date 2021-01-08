package com.atlas.cdc.processor;

import com.atlas.cdc.model.Character;
import com.atlas.cdc.model.Monster;
import com.atlas.cos.rest.attribute.CharacterAttributes;
import com.atlas.morg.rest.attribute.MonsterAttributes;

import rest.DataBody;

public final class ModelFactory {
   private ModelFactory() {
   }

   public static Character createCharacter(DataBody<CharacterAttributes> body) {
      return new Character(Integer.parseInt(body.getId()),
            body.getAttributes().jobId(),
            body.getAttributes().mapId(),
            body.getAttributes().x(),
            body.getAttributes().y(),
            body.getAttributes().hp(),
            body.getAttributes().mp(),
            body.getAttributes().meso()
      );
   }

   public static Monster createMonster(DataBody<MonsterAttributes> body) {
      return new Monster(Integer.parseInt(body.getId()),
            body.getAttributes().monsterId(),
            body.getAttributes().mapId()
      );
   }
}
