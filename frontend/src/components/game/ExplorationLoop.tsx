'use client';

import { motion, AnimatePresence } from 'framer-motion';
import { Node, StrategicChoice } from '@/services/exploration';
import { useState } from 'react';
import { explorationService } from '@/services/exploration';

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
  scrapMetal: number;
  researchData: number;
  encounters: Encounter[];
  timelineNodes: Node[];
  currentEncounter: Encounter | null;
  isTransitioning: boolean;
  onAdvance: () => void;
  onEnterCombat: (enemyId: string) => void;
  onResolveNode: (nodeId: string, choiceId?: string) => Promise<void>;
}

export default function ExplorationLoop({
  expeditionTitle,
  o2,
  fuel,
  scrapMetal,
  researchData,
  encounters,
  timelineNodes,
  currentEncounter,
  isTransitioning,
  onAdvance,
  onEnterCombat,
  onResolveNode
}: ExplorationLoopProps) {
  const [selectedNode, setSelectedNode] = useState<Node | null>(null);
  const [isResolving, setIsResolving] = useState(false);

  const handleResolveChoice = async (nodeId: string, choice: string) => {
    try {
      setIsResolving(true);
      await onResolveNode(nodeId, choice);
      setSelectedNode(null);
      onAdvance();
    } catch (error) {
      console.error('Failed to resolve choice:', error);
    } finally {
      setIsResolving(false);
    }
  };

  const handleMainAction = () => {
    if (isTransitioning || isResolving) return;

    if (!currentEncounter) {
      onAdvance();
      return;
    }

    if (currentEncounter.type === 'COMBAT') {
      onEnterCombat(currentEncounter.enemy_id || '');
    } else {
      // Check if the CURRENT node is resolved
      const currentNode = timelineNodes[encounters.length - 1];
      if (currentNode && !currentNode.is_resolved) {
        setSelectedNode(currentNode);
      } else {
        onAdvance();
      }
    }
  };

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
          <div className="text-right">
            <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Scrap Metal</div>
            <div className="text-xl font-black text-yellow-500">{scrapMetal}</div>
          </div>
          <div className="text-right">
            <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Research Data</div>
            <div className="text-xl font-black text-cyan-500">{researchData}</div>
          </div>
        </div>
      </div>

      {/* TIMELINE HUD */}
      <div className="px-6 py-4 bg-zinc-900/30 border-b border-zinc-900 flex items-center gap-4 overflow-x-auto no-scrollbar">
        {timelineNodes.map((node, idx) => (
          <div key={node.id} className="flex items-center gap-4 shrink-0">
            <div 
              onClick={() => !node.is_resolved && setSelectedNode(node)}
              className={`w-10 h-10 border flex items-center justify-center cursor-pointer transition-all ${
                node.is_resolved ? 'bg-zinc-800 border-zinc-700 opacity-50' : 
                idx === encounters.length - 1 ? 'border-pink-500 bg-pink-500/10 shadow-[0_0_10px_rgba(236,72,153,0.3)]' : 
                'border-zinc-800 hover:border-zinc-600'
              }`}
            >
              <span className="text-[10px] font-black">
                {node.type === 'COMBAT' ? '⚔️' : node.type === 'RESOURCE' ? '💎' : node.type === 'ANOMALY' ? '🌀' : '📍'}
              </span>
            </div>
            {idx < timelineNodes.length - 1 && (
              <div className="w-8 h-[1px] bg-zinc-800" />
            )}
          </div>
        ))}
      </div>

      {/* MAIN VIEWPORT */}
      <div className="flex-1 relative flex items-center justify-center p-4 md:p-12 overflow-hidden">
        <AnimatePresence mode="wait">
          {selectedNode ? (
            <motion.div 
              key="choice-overlay"
              initial={{ opacity: 0, scale: 0.95 }}
              animate={{ opacity: 1, scale: 1 }}
              exit={{ opacity: 0, scale: 0.95 }}
              className="z-20 w-full max-w-2xl bg-black/90 border border-zinc-800 p-8 backdrop-blur-xl"
            >
              <div className="text-[10px] text-pink-500 font-bold mb-2 tracking-[0.3em] uppercase">Strategic Encounter</div>
              <h2 className="text-2xl font-black italic mb-4">{(selectedNode.name || 'Unknown').toUpperCase()} // {(selectedNode.type || 'STORY').toUpperCase()}</h2>
              <p className="text-zinc-400 text-sm mb-8 leading-relaxed">{selectedNode.environment_description || 'No data available.'}</p>
              
              <div className="grid gap-4">
                {selectedNode.choices.map((choice) => (
                  <button
                    key={choice.label}
                    disabled={isResolving}
                    onClick={() => handleResolveChoice(selectedNode.id, choice.label)}
                    className="group relative border border-zinc-800 p-4 text-left hover:border-white transition-all disabled:opacity-50"
                  >
                    <div className="flex justify-between items-start mb-2">
                      <span className="font-black text-sm uppercase tracking-widest">{choice.label}</span>
                      <span className="text-[10px] text-zinc-500">Success: {Math.round(choice.success_chance * 100)}%</span>
                    </div>
                    <p className="text-xs text-zinc-500 group-hover:text-zinc-300 transition-colors">{choice.description}</p>
                    {choice.requirements.length > 0 && (
                      <div className="mt-2 flex gap-2">
                        {choice.requirements.map(req => (
                          <span key={req} className="text-[8px] bg-zinc-900 px-2 py-1 text-zinc-400 border border-zinc-800">{req}</span>
                        ))}
                      </div>
                    )}
                  </button>
                ))}
              </div>
              
              <button 
                onClick={() => setSelectedNode(null)}
                className="mt-8 text-[10px] text-zinc-500 hover:text-white uppercase tracking-widest"
              >
                Cancel / Re-evaluate
              </button>
            </motion.div>
          ) : currentEncounter ? (
            <motion.div 
              key={currentEncounter.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: -10 }}
              className="relative w-full max-w-4xl aspect-video border border-zinc-800 overflow-hidden group cursor-pointer"
              onClick={handleMainAction}
            >
              <img src={currentEncounter.image} className="w-full h-full object-cover opacity-60 grayscale group-hover:grayscale-0 transition-all duration-700" />
              <div className="absolute inset-0 bg-gradient-to-t from-black via-transparent to-transparent" />
              
              {/* CLICK TO ADVANCE HINT */}
              <div className="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                <div className="bg-white/10 backdrop-blur-md border border-white/20 px-6 py-3 uppercase text-[10px] tracking-[0.3em] font-black">
                  {currentEncounter.type === 'COMBAT' ? 'Initiate Combat' : 'Analyze / Advance'}
                </div>
              </div>

              {/* DNA OVERLAY */}
              <div className="absolute bottom-4 left-4 right-4 md:bottom-8 md:left-8 md:right-8">
                <div className="text-[10px] text-pink-500 font-bold mb-2 tracking-[0.3em] uppercase">Visual DNA Synthesis</div>
                <div className="text-[10px] md:text-xs text-zinc-400 font-mono bg-black/60 p-2 md:p-3 border-l-2 border-pink-500 backdrop-blur-sm">
                  <span className="text-zinc-500">DNA_KEYWORDS:</span> {currentEncounter.visualPrompt?.toUpperCase() || 'SCANNING...'}
                </div>
              </div>

              {/* ENCOUNTER INFO */}
              <div className="absolute top-4 left-4 md:top-8 md:left-8">
                <div className="bg-black/80 border border-zinc-800 p-4 backdrop-blur-md">
                  <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">{currentEncounter.type || 'STORY'}</div>
                  <div className="text-xl font-black italic uppercase tracking-tighter">{currentEncounter.title || 'Unknown Encounter'}</div>
                  <p className="text-xs text-zinc-400 mt-2 max-w-xs leading-relaxed">{currentEncounter.description || 'Scanning environment...'}</p>
                </div>
              </div>
            </motion.div>
          ) : (
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              className="text-center"
            >
              <div className="text-pink-500 text-xl mb-4 animate-pulse font-black tracking-widest">INITIALIZING EXPEDITION...</div>
              <button 
                onClick={onAdvance}
                className="px-12 py-4 bg-white text-black font-black hover:bg-pink-500 hover:text-white transition-all uppercase tracking-widest border border-white"
              >
                Begin Descent
              </button>
            </motion.div>
          )}
        </AnimatePresence>

        {/* TRANSITION OVERLAY */}
        {isTransitioning && (
          <motion.div 
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            className="absolute inset-0 z-50 bg-black/60 backdrop-blur-sm flex items-center justify-center"
          >
            <div className="flex flex-col items-center">
              <div className="w-12 h-12 border-2 border-pink-500 border-t-transparent rounded-full animate-spin mb-4" />
              <div className="text-[10px] font-black tracking-[0.5em] text-pink-500 animate-pulse uppercase">Synchronizing Timeline...</div>
            </div>
          </motion.div>
        )}
      </div>

      {/* FOOTER ACTION */}
      <div className="p-6 border-t border-zinc-900 bg-zinc-950/50 backdrop-blur-xl flex justify-between items-center">
        <div className="flex flex-col">
          <div className="text-[10px] text-zinc-600 uppercase tracking-widest">
            Step {encounters.length} / Narrative Timeline
          </div>
          <div className="text-[8px] text-zinc-500 uppercase tracking-widest mt-1">
            Nodes Discovered: {encounters.length} / {timelineNodes.length}
          </div>
        </div>

        <button 
          onClick={handleMainAction}
          disabled={o2 <= 0 || isTransitioning}
          className="bg-white text-black px-12 py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20 relative group"
        >
          <div className="flex flex-col items-center">
            <span>{currentEncounter?.type === 'COMBAT' ? 'Enter Combat' : 'Advance Timeline'}</span>
            {currentEncounter?.type !== 'COMBAT' && (
              <span className="text-[8px] font-bold opacity-60 group-hover:opacity-100">
                COST: 15 O2 | 5 FUEL
              </span>
            )}
          </div>
        </button>
      </div>
    </div>
  );
}
