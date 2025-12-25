'use client';

import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { PullType } from '@/services/gacha';
import { gachaSystem, GachaState } from '@/systems/GachaSystem';
import { gameEvents, GAME_EVENTS } from '@/systems/EventBus';

interface GachaStageProps {
  onBack: () => void;
}

export default function GachaStage({ onBack }: GachaStageProps) {
  const [gachaState, setGachaState] = useState<GachaState>(gachaSystem.getState());

  useEffect(() => {
    const unsubscribe = gameEvents.on(GAME_EVENTS.GACHA_PULLED, (newState: GachaState) => {
      setGachaState(newState);
    });

    return () => {
      unsubscribe();
      gachaSystem.clearResults();
    };
  }, []);

  const { isPulling, results, error } = gachaState;

  const handlePull = async (type: PullType) => {
    await gachaSystem.pull(type);
  };

  return (
    <div className="h-full w-full flex flex-col bg-black text-white font-mono p-8 overflow-y-auto">
      <div className="flex justify-between items-center mb-12">
        <h2 className="text-4xl font-black italic text-yellow-400 tracking-tighter">VOID SIGNALS</h2>
        <button 
          onClick={onBack}
          className="text-zinc-500 hover:text-white transition-colors uppercase text-xs tracking-widest"
        >
          [ Back to Bastion ]
        </button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
        {/* Daily Signal */}
        <div className="border border-zinc-800 p-6 bg-zinc-900/20 flex flex-col items-center text-center">
          <div className="w-20 h-20 bg-zinc-800 rounded-full mb-4 flex items-center justify-center border border-zinc-700">
            <div className="w-12 h-12 bg-yellow-400/20 rounded-full animate-pulse" />
          </div>
          <h3 className="text-xl font-bold mb-2">Daily Signal</h3>
          <p className="text-xs text-zinc-500 mb-6">A faint echo from the void. Available every 24 hours.</p>
          <button 
            onClick={() => handlePull('DAILY_SIGNAL')}
            disabled={isPulling}
            className="w-full py-2 bg-yellow-400 text-black font-bold uppercase tracking-widest hover:bg-yellow-300 disabled:opacity-50 transition-all"
          >
            {isPulling ? 'Scanning...' : 'Intercept'}
          </button>
        </div>

        {/* Relic Signal */}
        <div className="border border-zinc-800 p-6 bg-zinc-900/20 flex flex-col items-center text-center">
          <div className="w-20 h-20 bg-zinc-800 rounded-full mb-4 flex items-center justify-center border border-zinc-700">
            <div className="w-12 h-12 bg-blue-400/20 rounded-full" />
          </div>
          <h3 className="text-xl font-bold mb-2">Relic Signal</h3>
          <p className="text-xs text-zinc-500 mb-6">Guaranteed Rare or higher every 10 pulls.</p>
          <button 
            onClick={() => handlePull('RELIC_SIGNAL')}
            disabled={isPulling}
            className="w-full py-2 border border-blue-400 text-blue-400 font-bold uppercase tracking-widest hover:bg-blue-400 hover:text-black disabled:opacity-50 transition-all"
          >
            {isPulling ? 'Scanning...' : '160 Credits'}
          </button>
        </div>

        {/* Singularity Signal */}
        <div className="border border-zinc-800 p-6 bg-zinc-900/20 flex flex-col items-center text-center">
          <div className="w-20 h-20 bg-zinc-800 rounded-full mb-4 flex items-center justify-center border border-zinc-700">
            <div className="w-12 h-12 bg-purple-400/20 rounded-full" />
          </div>
          <h3 className="text-xl font-bold mb-2">Singularity</h3>
          <p className="text-xs text-zinc-500 mb-6">High chance for Void-Touched parts.</p>
          <button 
            onClick={() => handlePull('SINGULARITY_SIGNAL')}
            disabled={isPulling}
            className="w-full py-2 border border-purple-400 text-purple-400 font-bold uppercase tracking-widest hover:bg-purple-400 hover:text-black disabled:opacity-50 transition-all"
          >
            {isPulling ? 'Scanning...' : '320 Credits'}
          </button>
        </div>
      </div>

      <AnimatePresence>
        {error && (
          <motion.div 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0 }}
            className="p-4 bg-red-900/20 border border-red-500 text-red-400 text-sm text-center mb-8"
          >
            {error}
          </motion.div>
        )}

        {results && (
          <motion.div 
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            className="grid grid-cols-2 md:grid-cols-5 gap-4"
          >
            {results.map((item, i) => (
              <motion.div 
                key={i}
                initial={{ scale: 0.8, opacity: 0 }}
                animate={{ scale: 1, opacity: 1 }}
                transition={{ delay: i * 0.1 }}
                className="border border-zinc-800 p-4 bg-zinc-900/40 text-center"
              >
                <div className="text-[10px] text-zinc-500 uppercase mb-1">{item.rarity}</div>
                <div className="font-bold text-sm mb-2">{item.name || 'Unknown Part'}</div>
                <div className="text-[10px] text-zinc-400">{item.slot}</div>
              </motion.div>
            ))}
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}
