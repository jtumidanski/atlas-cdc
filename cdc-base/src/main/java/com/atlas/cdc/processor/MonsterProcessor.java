package com.atlas.cdc.processor;

import com.app.rest.util.RestResponseUtil;
import com.atlas.cdc.model.LoseItem;
import com.atlas.cdc.model.Monster;
import com.atlas.cdc.model.MonsterInformation;
import com.atlas.morg.rest.attribute.MonsterAttributes;
import com.atlas.morg.rest.constant.RestConstants;
import com.atlas.shared.rest.UriBuilder;

import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.CompletableFuture;

public final class MonsterProcessor {
   private MonsterProcessor() {
   }

   public static CompletableFuture<Optional<Monster>> getFromUniqueId(int uniqueId) {
      return UriBuilder.service(RestConstants.SERVICE)
            .pathParam("monsters", uniqueId)
            .getAsyncRestClient(MonsterAttributes.class)
            .get()
            .thenApply(RestResponseUtil.data(ModelFactory::createMonster))
            .thenApply(monster -> monster.map(MonsterProcessor::associateInformation));
   }

   protected static Monster associateInformation(Monster monster) {
      MonsterInformation information = MonsterInformationProcessor.getById(monster.id()).join();
      return monster.setInformation(information);
   }

   public static List<LoseItem> getLoseItemList(int monsterId) {
      return Collections.emptyList();
   }

   public static boolean isBuffed(int uniqueId, Object buff) {
      return false;
   }

   public static boolean isNeutralized(int uniqueId) {
      return isBuffed(uniqueId, "MonsterStatus.NEUTRALISE");
   }
}
