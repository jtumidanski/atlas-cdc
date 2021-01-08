package com.atlas.cdc.processor;

import java.util.concurrent.CompletableFuture;

import com.atlas.cdc.model.MonsterInformation;

public final class MonsterInformationProcessor {
   private MonsterInformationProcessor() {
   }

   public static CompletableFuture<MonsterInformation> getById(int monsterId) {
      return CompletableFuture.completedFuture(new MonsterInformation(0));
   }
}
