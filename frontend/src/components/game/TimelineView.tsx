'use client';

import { motion, AnimatePresence } from 'framer-motion';
import { Node, StrategicChoice } from '@/services/exploration';

interface TimelineViewProps {
  expeditionTitle: string;
  o2: number;
  fuel: number;
  scrap?: number;
  research?: number;
  timeline: Node[];
  currentNode: Node | null;
  onResolveChoice: (nodeId: string, choice: string) => void;
  onEnterCombat: (enemyId: string) => void;
  onAdvance: () => void;
}

export default function TimelineView({
  expeditionTitle,
  o2,
  fuel,
  scrap,
  research,
  timeline,
  currentNode,
  onResolveChoice,
  onEnterCombat,
  onAdvance
}: TimelineViewProps) {
  return (
    <div className="h-screen flex flex-col bg-black text-white font-mono overflow-hidden">
      {/* HUD TOP */}
      <div className="p-6 flex justify-between items-start border-b border-zinc-900 bg-black/80 backdrop-blur-md z-10">
        <div>
          <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Current Expedition</div>
          <div className="text-lg font-black italic text-pink-500">{expeditionTitle}</div>
        </div>
        <div className="flex gap-8">
          <div className="text-right">
            <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Scrap</div>
            <div className="text-xl font-black text-yellow-500">{scrap || 0}</div>
          </div>
          <div className="text-right">
            <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Research</div>
            <div className="text-xl font-black text-blue-500">{research || 0}</div>
          </div>
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

      {/* TIMELINE TRACKER */}
      <div className="px-6 py-4 bg-zinc-950 border-b border-zinc-900 flex items-center gap-4 overflow-x-auto no-scrollbar">
        {timeline.map((node, idx) => (
          <div key={node.id} className="flex items-center gap-4 shrink-0">
            <div className={`flex flex-col items-center ${node.id === currentNode?.id ? 'opacity-100' : 'opacity-40'}`}>
              <div className={`w-8 h-8 rounded-full border flex items-center justify-center text-[10px] font-bold
                ${node.is_resolved ? 'bg-pink-500 border-pink-500 text-black' : 'border-zinc-700'}
                ${node.id === currentNode?.id ? 'ring-2 ring-white ring-offset-2 ring-offset-black' : ''}
              `}>
                {idx + 1}
              </div>
              <div className="text-[8px] mt-1 uppercase tracking-tighter">{node.type}</div>
            </div>
            {idx < timeline.length - 1 && (
              <div className="w-12 h-[1px] bg-zinc-800" />
            )}
          </div>
        ))}
      </div>

      {/* MAIN VIEWPORT */}
      <div className="flex-1 relative flex flex-col md:flex-row">
        {/* VISUAL DNA AREA */}
        <div className="flex-1 relative border-r border-zinc-900 bg-zinc-950 flex items-center justify-center p-8">
          <AnimatePresence mode="wait">
            <motion.div 
              key={currentNode?.id}
              initial={{ opacity: 0, scale: 0.95 }}
              animate={{ opacity: 1, scale: 1 }}
              exit={{ opacity: 0, scale: 1.05 }}
              className="relative w-full max-w-2xl aspect-square border border-zinc-800 overflow-hidden group"
            >
              <img 
                src="https://images.unsplash.com/photo-1614728263952-84ea256f9679?q=80&w=1000&auto=format&fit=crop" 
                className="w-full h-full object-cover opacity-40 grayscale group-hover:grayscale-0 transition-all duration-1000" 
              />
              <div className="absolute inset-0 bg-gradient-to-t from-black via-transparent to-transparent" />
              
              {/* SCANLINE EFFECT */}
              <div className="absolute inset-0 pointer-events-none bg-[linear-gradient(rgba(18,16,16,0)_50%,rgba(0,0,0,0.25)_50%),linear-gradient(90deg,rgba(255,0,0,0.06),rgba(0,255,0,0.02),rgba(0,0,255,0.06))] bg-[length:100%_2px,3px_100%] z-10" />

              <div className="absolute bottom-8 left-8 right-8">
                <div className="text-[10px] text-pink-500 font-bold mb-2 tracking-[0.3em] uppercase">Visual DNA Synthesis</div>
                <div className="text-xs text-zinc-400 font-mono bg-black/60 p-4 border-l-2 border-pink-500 backdrop-blur-sm">
                  {currentNode?.environment_description?.toUpperCase() || 'SCANNING ENVIRONMENT...'}
                </div>
              </div>
            </motion.div>
          </AnimatePresence>
        </div>

        {/* INTERACTION AREA */}
        <div className="w-full md:w-[400px] bg-black p-8 flex flex-col">
          <div className="mb-8">
            <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-2">Node Encounter</div>
            <h2 className="text-2xl font-black italic mb-4">{currentNode?.name}</h2>
            <p className="text-sm text-zinc-400 leading-relaxed">
              {currentNode?.type === 'COMBAT' ? 'Hostile signatures detected. Prepare for engagement.' : 
               currentNode?.type === 'RESOURCE' ? 'Valuable mineral deposits identified in the vicinity.' :
               'Navigating through the silent void. No immediate threats detected.'}
            </p>
          </div>

          <div className="flex-1 space-y-4">
            <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-4">Strategic Choices</div>
            {currentNode?.choices?.map((choice) => (
              <button
                key={choice.label}
                disabled={currentNode?.is_resolved}
                onClick={() => currentNode && onResolveChoice(currentNode.id, choice.label)}
                className={`w-full text-left p-4 border transition-all group relative overflow-hidden
                  ${currentNode?.is_resolved ? 'border-zinc-800 opacity-50 cursor-not-allowed' : 'border-zinc-800 hover:border-pink-500 hover:bg-pink-500/5'}
                `}
              >
                <div className="flex justify-between items-start mb-1">
                  <span className="font-bold text-sm uppercase tracking-wider group-hover:text-pink-500">{choice.label}</span>
                  <span className="text-[10px] text-zinc-500">{(choice.success_chance * 100).toFixed(0)}% SUCCESS</span>
                </div>
                <p className="text-[10px] text-zinc-500 leading-tight">{choice.description}</p>
                
                {choice.requirements && choice.requirements.length > 0 && (
                  <div className="mt-2 flex gap-2">
                    {choice.requirements.map(req => (
                      <span key={req} className="text-[8px] bg-zinc-900 text-zinc-400 px-1.5 py-0.5 rounded border border-zinc-800">
                        REQ: {req}
                      </span>
                    ))}
                  </div>
                )}

                {/* HOVER DECOR */}
                <div className="absolute top-0 right-0 w-1 h-0 bg-pink-500 transition-all group-hover:h-full" />
              </button>
            ))}

            {currentNode?.is_resolved && (
              <motion.button 
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                onClick={onAdvance}
                className="w-full p-6 bg-white text-black font-black uppercase tracking-widest hover:bg-pink-500 hover:text-white transition-all shadow-[0_0_20px_rgba(255,255,255,0.1)]"
              >
                <div className="text-[10px] opacity-50 mb-1">Node Resolved</div>
                <div className="text-sm">Proceed to Next Sector</div>
              </motion.button>
            )}
          </div>

          <div className="mt-8 pt-8 border-t border-zinc-900">
            <div className="flex justify-between text-[10px] text-zinc-500 mb-4">
              <span>SYSTEM STATUS</span>
              <span className="text-green-500">ONLINE</span>
            </div>
            <div className="grid grid-cols-4 gap-1">
              {[...Array(12)].map((_, i) => (
                <div key={i} className={`h-1 ${i < 8 ? 'bg-pink-500/20' : 'bg-zinc-900'}`} />
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
