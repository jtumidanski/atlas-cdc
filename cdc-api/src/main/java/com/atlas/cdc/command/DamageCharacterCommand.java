package com.atlas.cdc.command;

public record DamageCharacterCommand(int characterId, int monsterId, int monsterUniqueId, byte damageFrom, byte element,
                                     int damage, byte direction) {
}
