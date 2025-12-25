'use client';

import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { DamageType } from '@/services/combat';
import { combatSystem, CombatState } from '@/systems/CombatSystem';
import { gameEvents, GAME_EVENTS } from '@/systems/EventBus';

interface CombatStageProps {
  attackerId: string;
  enemyId: string;
  onCombatEnd: (result: 'VICTORY' | 'DEFEAT' | 'ESCAPE') => void;
}

export default function CombatStage({
  attackerId,
  enemyId,
  onCombatEnd
}: CombatStageProps) {
  const [combatState, setCombatState] = useState<CombatState>(combatSystem.getState());

  useEffect(() => {
    const unsubscribe = gameEvents.on(GAME_EVENTS.COMBAT_UPDATED, (newState: CombatState) => {
      setCombatState(newState);
    });

    const endUnsubscribe = gameEvents.on(GAME_EVENTS.COMBAT_ENDED, ({ result }) => {
      setTimeout(() => onCombatEnd(result), 2000);
    });

    // Initialize combat if not already started
    combatSystem.startCombat(attackerId, enemyId);

    return () => {
      unsubscribe();
      endUnsubscribe();
    };
  }, [attackerId, enemyId, onCombatEnd]);

  const { attackerStats, defenderStats, combatLog, isProcessing, turn } = combatState;

  const attackerName = "PLAYER_VEHICLE";
  const defenderName = "ENEMY_UNIT";

  const handleAttack = async (type: DamageType) => {
    await combatSystem.executeAttack(attackerId, enemyId, type);
  };

  return (
    <div className="h-full w-full flex flex-col bg-black text-white font-mono p-8">
      <div className="flex justify-between items-center mb-12">
        <h2 className="text-4xl font-black italic text-red-500 tracking-tighter">COMBAT ENGAGED</h2>
        <div className="text-zinc-500 text-xs uppercase tracking-widest">Turn {turn}</div>
      </div>

      <div className="flex-1 grid grid-cols-1 md:grid-cols-2 gap-12 mb-8">
        {/* Player Side */}
        <div className="border border-zinc-800 p-8 bg-zinc-900/20 relative overflow-hidden">
          <div className="absolute top-0 left-0 w-1 h-full bg-blue-500" />
          <div className="text-xs text-zinc-500 uppercase mb-2">Friendly Unit</div>
          <div className="text-3xl font-bold mb-6">{attackerName}</div>
          
          <div className="space-y-4">
            <div>
              <div className="flex justify-between text-xs mb-1">
                <span>INTEGRITY</span>
                <span>{attackerStats?.hp || '???'}/{attackerStats?.max_hp || '???'}</span>
              </div>
              <div className="w-full h-2 bg-zinc-900">
                <motion.div 
                  className="h-full bg-blue-500"
                  initial={{ width: '100%' }}
                  animate={{ width: attackerStats ? `${(attackerStats.hp / attackerStats.max_hp) * 100}%` : '100%' }}
                />
              </div>
            </div>
            
            <div className="grid grid-cols-2 gap-4 text-[10px] text-zinc-400 uppercase">
              <div className="border border-zinc-800 p-2">ATK: {attackerStats?.base_attack || '--'}</div>
              <div className="border border-zinc-800 p-2">DEF: {attackerStats?.target_defense || '--'}</div>
            </div>
          </div>
        </div>

        {/* Enemy Side */}
        <div className="border border-zinc-800 p-8 bg-zinc-900/20 relative overflow-hidden">
          <div className="absolute top-0 right-0 w-1 h-full bg-red-500" />
          <div className="text-xs text-zinc-500 uppercase mb-2 text-right">Hostile Signature</div>
          <div className="text-3xl font-bold mb-6 text-right">{defenderName}</div>
          
          <div className="space-y-4">
            <div>
              <div className="flex justify-between text-xs mb-1">
                <span>INTEGRITY</span>
                <span>{defenderStats?.hp || '???'}/{defenderStats?.max_hp || '???'}</span>
              </div>
              <div className="w-full h-2 bg-zinc-900">
                <motion.div 
                  className="h-full bg-red-500"
                  initial={{ width: '100%' }}
                  animate={{ width: defenderStats ? `${(defenderStats.hp / defenderStats.max_hp) * 100}%` : '100%' }}
                />
              </div>
            </div>

            <div className="grid grid-cols-2 gap-4 text-[10px] text-zinc-400 uppercase">
              <div className="border border-zinc-800 p-2">ATK: {defenderStats?.base_attack || '--'}</div>
              <div className="border border-zinc-800 p-2">DEF: {defenderStats?.target_defense || '--'}</div>
            </div>
          </div>
        </div>
      </div>

      {/* Combat Log & Controls */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-8 h-48">
        <div className="md:col-span-2 border border-zinc-800 bg-black p-4 overflow-y-auto text-xs space-y-1 font-mono">
          <AnimatePresence>
            {combatLog.map((log, i) => (
              <motion.div 
                key={`log-${i}`}
                initial={{ opacity: 0, x: -10 }}
                animate={{ opacity: 1, x: 0 }}
                className={log.includes('Error') ? 'text-red-500' : log.includes('CRITICAL') ? 'text-yellow-500 font-bold' : 'text-zinc-400'}
              >
                {`> ${log}`}
              </motion.div>
            ))}
          </AnimatePresence>
          {combatLog.length === 0 && <div className="text-zinc-700 italic">Waiting for tactical input...</div>}
        </div>

        <div className="flex flex-col gap-2">
          <button 
            onClick={() => handleAttack('KINETIC')}
            disabled={isProcessing}
            className="flex-1 bg-white text-black font-black uppercase tracking-tighter hover:bg-blue-500 hover:text-white transition-all disabled:opacity-50"
          >
            Kinetic Strike
          </button>
          <button 
            onClick={() => handleAttack('ENERGY')}
            disabled={isProcessing}
            className="flex-1 border border-zinc-700 text-zinc-300 font-black uppercase tracking-tighter hover:bg-purple-500 hover:text-white transition-all disabled:opacity-50"
          >
            Energy Pulse
          </button>
          <button 
            onClick={() => onCombatEnd('ESCAPE')}
            disabled={isProcessing}
            className="py-2 text-[10px] text-zinc-500 uppercase tracking-widest hover:text-white transition-colors"
          >
            Attempt Emergency Retreat
          </button>
        </div>
      </div>
    </div>
  );
}
