package com.atlas.cdc.processor;

import java.util.Optional;

import com.atlas.cdc.model.Skill;
import com.atlas.cdc.model.SkillEffect;

public final class SkillProcessor {
   private SkillProcessor() {
   }

   public static Optional<Skill> getSkillById(int skillId) {
      return Optional.empty();
   }

   public static int getSkillLevel(int characterId, int skillId) {
      return 0;
   }

   public static Optional<SkillEffect> getSkillEffect(int skillId, int skillLevel) {
      return Optional.empty();
   }

   public static Optional<SkillEffect> getCharacterSkillEffect(int characterId, int skillId) {
      return Optional.empty();
   }
}
