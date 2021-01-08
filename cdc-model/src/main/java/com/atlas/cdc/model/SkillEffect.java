package com.atlas.cdc.model;

public record SkillEffect(int x) {
   public boolean makeChanceResult() {
      return false;
   }
}
