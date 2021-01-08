package com.atlas.cdc.model;

public record Monster(int id, int monsterId, int mapId, int maxHp) {
   public Monster(int id, int monsterId, int mapId) {
      this(id, monsterId, mapId, 0);
   }

   public Monster setInformation(MonsterInformation information) {
      return new Monster(id, monsterId, mapId, information.maxHp());
   }

   public boolean boss() {
      return false;
   }
}
