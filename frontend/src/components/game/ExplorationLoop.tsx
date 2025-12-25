'use client';

import { motion, AnimatePresence } from 'framer-motion';

interface Encounter {
  id: string;
  type: string;
  title: string;
  description: string;
  visualPrompt: string;
  image: string;
  enemy_id?: string;
}

interface ExplorationLoopProps {
  expeditionTitle: string;
  o2: number;
  fuel: number;
  encounters: Encounter[];
  currentEncounter: Encounter | null;
  isTransitioning: boolean;
  onAdvance: () => void;
  onEnterCombat: (enemyId: string) => void;
}

export default function ExplorationLoop({
  expeditionTitle,
  o2,
  fuel,
  encounters,
  currentEncounter,
  isTransitioning,
  onAdvance,
  onEnterCombat
}: ExplorationLoopProps) {
  return (
    <div className="h-screen flex flex-col bg-black text-white font-mono">
      {/* HUD TOP */}
      <div className="p-6 flex justify-between items-start border-b border-zinc-900 bg-black/80 backdrop-blur-md z-10">
        <div>
          <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Current Expedition</div>
          <div className="text-lg font-black italic text-pink-500">{expeditionTitle}</div>
        </div>
        <div className="flex gap-8">
          <div className="text-right">
            <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Oxygen Level</div>
            <div className={`text-xl font-black ${o2 < 30 ? 'text-red-500 animate-pulse' : 'text-white'}`}>{o2}%</div>
            <div className="w-32 h-1 bg-zinc-900 mt-1">
              <motion.div className="h-full bg-white" animate={{ width: `${o2}%` }} />
            </div>
          </div>
          <div className="text-right">
            <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Fuel Reserve</div>
            <div className="text-xl font-black">{fuel}%</div>
            <div className="w-32 h-1 bg-zinc-900 mt-1">
              <motion.div className="h-full bg-zinc-500" animate={{ width: `${fuel}%` }} />
            </div>
          </div>
        </div>
      </div>

      {/* MAIN VIEWPORT */}
      <div 
        className="flex-1 relative flex items-center justify-center p-4 md:p-12 cursor-pointer group"
        onClick={() => {
          if (currentEncounter?.type === 'COMBAT') {
            onEnterCombat(currentEncounter.enemy_id || '');
          } else {
            onAdvance();
          }
        }}
      >
        <AnimatePresence mode="wait">
          {currentEncounter && (
            <motion.div 
              key={currentEncounter.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: -10 }}
              className="relative w-full max-w-4xl aspect-video border border-zinc-800 overflow-hidden"
            >
              <img src={currentEncounter.image} className="w-full h-full object-cover opacity-60 grayscale group-hover:grayscale-0 transition-all duration-700" />
              <div className="absolute inset-0 bg-gradient-to-t from-black via-transparent to-transparent" />
              
              {/* CLICK TO ADVANCE HINT */}
              <div className="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                <div className="bg-white/10 backdrop-blur-md border border-white/20 px-6 py-3 uppercase text-[10px] tracking-[0.3em] font-black">
                  Click to Advance
                </div>
              </div>

              {/* DNA OVERLAY */}
              <div className="absolute bottom-4 left-4 right-4 md:bottom-8 md:left-8 md:right-8">
                <div className="text-[10px] text-pink-500 font-bold mb-2 tracking-[0.3em] uppercase">Visual DNA Synthesis</div>
                <div className="text-[10px] md:text-xs text-zinc-400 font-mono bg-black/60 p-2 md:p-3 border-l-2 border-pink-500 backdrop-blur-sm">
                  <span className="text-zinc-500">DNA_KEYWORDS:</span> {currentEncounter.visualPrompt.toUpperCase()}
                </div>
              </div>

              {/* ENCOUNTER INFO */}
              <div className="absolute top-4 left-4 md:top-8 md:left-8">
                <div className="bg-white text-black px-2 py-0.5 md:px-3 md:py-1 text-[8px] md:text-[10px] font-black uppercase mb-1 md:mb-2 inline-block">
                  {currentEncounter.type}
                </div>
                <h3 className="text-xl md:text-3xl font-black italic tracking-tighter">{currentEncounter.title}</h3>
                <p className="text-zinc-400 max-w-xs md:max-w-md mt-1 md:mt-2 text-[10px] md:text-sm leading-relaxed">{currentEncounter.description}</p>
              </div>
            </motion.div>
          )}
        </AnimatePresence>
      </div>

      {/* TIMELINE STRING */}
      <div className="p-4 md:p-8 border-t border-zinc-900 bg-zinc-950/50 backdrop-blur-xl">
        <div className="flex items-center gap-4 mb-6 overflow-x-auto pb-4 no-scrollbar">
          {encounters.map((e, i) => (
            <div key={`timeline-${e.id}-${i}`} className="flex items-center gap-4 shrink-0">
              <div className={`w-3 h-3 rounded-full ${
                e.type === 'COMBAT' ? 'bg-red-500' : 
                e.type === 'RESOURCE' ? 'bg-blue-500' : 
                'bg-white'
              } ${i === encounters.length - 1 ? 'ring-4 ring-pink-500/20 scale-125' : 'opacity-40'}`} />
              {i < encounters.length - 1 && <div className="w-12 h-[1px] bg-zinc-800" />}
            </div>
          ))}
          <div className="w-12 h-[1px] bg-zinc-800 border-dashed border-t" />
          <div className="w-3 h-3 rounded-full border border-zinc-800" />
        </div>

        <div className="flex justify-between items-center">
          <div className="flex flex-col">
            <div className="text-[10px] text-zinc-600 uppercase tracking-widest">
              Step {encounters.length} / Narrative Timeline
            </div>
            {isTransitioning && (
              <div className="text-[10px] text-red-500 font-bold animate-pulse uppercase mt-1">
                Target Lock Detected - Transitioning to Combat...
              </div>
            )}
          </div>
          <button 
            onClick={() => {
              if (currentEncounter?.type === 'COMBAT') {
                onEnterCombat(currentEncounter.enemy_id || '');
              } else {
                onAdvance();
              }
            }}
            disabled={o2 <= 0 || isTransitioning}
            className="bg-white text-black px-12 py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20 relative group"
          >
            <div className="flex flex-col items-center">
              <span>{currentEncounter?.type === 'COMBAT' ? 'Enter Combat' : (o2 < 30 ? 'Desperate Move' : 'Advance Timeline')}</span>
              {currentEncounter?.type !== 'COMBAT' && (
                <span className="text-[8px] font-bold opacity-60 group-hover:opacity-100">
                  COST: 15 O2 | 5 FUEL
                </span>
              )}
            </div>
          </button>
        </div>
      </div>
    </div>
  );
}
