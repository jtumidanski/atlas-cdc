package com.atlas.cdc.processor;

public final class CharacterSkillProcessor {
   private CharacterSkillProcessor() {
   }

   public static boolean hasAura(int characterId) {
      return isBuffed(characterId, "MapleBuffStat.AURA");
   }

   public static boolean isBuffed(int characterId, Object buff) {
      return buffValue(characterId, buff) != 0;
   }

   public static int buffValue(int characterId, Object buff) {
      return 0;
   }

   public static boolean hasPowerGuard(int characterId) {
      return isBuffed(characterId, "MapleBuffStat.POWER_GUARD");
   }

   public static boolean isBuffedFrom(int characterId, int skillId, Object buff) {
      return false;
   }

   public static boolean hasBodyPressure(int characterId) {
      return isBuffed(characterId, "MapleBuffStat.BODY_PRESSURE");
   }

   public static boolean hasMesoGuard(int characterId) {
      return isBuffed(characterId, "MapleBuffStat.MESO_GUARD");
   }

   public static boolean hasMagicGuard(int characterId) {
      return isBuffed(characterId, "MapleBuffStat.MAGIC_GUARD");
   }

   public static void cancelBuff(int characterId, Object buff) {
   }
}
