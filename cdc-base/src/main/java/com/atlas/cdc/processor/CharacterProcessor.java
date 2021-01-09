package com.atlas.cdc.processor;

import com.app.rest.util.RestResponseUtil;
import com.atlas.cdc.event.producer.AdjustHealthProducer;
import com.atlas.cdc.event.producer.AdjustManaProducer;
import com.atlas.cdc.event.producer.AdjustMesoProducer;
import com.atlas.cdc.model.Character;
import com.atlas.cos.constant.RestConstants;
import com.atlas.cos.rest.attribute.CharacterAttributes;
import com.atlas.shared.rest.UriBuilder;

import java.util.Optional;
import java.util.concurrent.CompletableFuture;

public final class CharacterProcessor {
   private CharacterProcessor() {
   }

   public static CompletableFuture<Optional<Character>> getFromId(int characterId) {
      return UriBuilder.service(RestConstants.SERVICE)
            .pathParam("characters", characterId)
            .getAsyncRestClient(CharacterAttributes.class)
            .get()
            .thenApply(RestResponseUtil.data(ModelFactory::createCharacter));
   }

   public static void removeItemFromInventory(int characterId, Object type, int itemId, int qty) {
      //MapleInventoryManipulator.removeById(client, type, loseItem.itemId(), qty, false, false);
   }

   public static int getAmountToDrop(int characterId, int itemId, int maxDroppable) {
      //Math.min(chr.countItem(loseItem.itemId()), maxDroppable)
      return 0;
   }

   public static void updateAriantScore(int characterId) {
   }

   public static boolean isRidingBattleship(int characterId) {
      return false;
   }

   public static void adjustBattleshipHp(int characterId, int amount) {
   }

   public static void adjustMeso(int characterId, int amount, boolean show) {
      AdjustMesoProducer.emit(characterId, amount, show);
   }

   public static void adjustMana(int characterId, int amount) {
      AdjustManaProducer.emit(characterId, amount);
   }

   public static void adjustHealth(int characterId, int amount) {
      AdjustHealthProducer.emit(characterId, amount);
   }

   public static boolean isHidden(int characterId) {
      return false;
   }
}
