import React, { useState, useEffect } from 'react';
import { useActor } from '@xstate/react';
import { GameContext } from '@/app/page';

interface ResearchNode {
  id: string;
  name: string;
  description: string;
  cost: number;
  unlocked: boolean;
}

const RESEARCH_NODES: ResearchNode[] = [
  { id: 'atmosphericEntry', name: 'Atmospheric Entry', description: 'Allows landing on planets with atmosphere.', cost: 100, unlocked: false },
  { id: 'miningDrill', name: 'Mining Drill', description: 'Extracts more resources from asteroids.', cost: 200, unlocked: false },
  { id: 'hackingModule', name: 'Hacking Module', description: 'Bypass security on derelict stations.', cost: 300, unlocked: false },
  { id: 'quantumGate', name: 'Quantum Gate', description: 'Travel to distant sectors instantly.', cost: 500, unlocked: false },
];

export const ResearchTerminal = () => {
  const gameService = React.useContext(GameContext);
  const [state, send] = useActor(gameService);
  const [nodes, setNodes] = useState<ResearchNode[]>(RESEARCH_NODES);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const { research, pilotStats } = state.context;
  const unlockedResearch = React.useMemo(() => 
    pilotStats?.metadata?.unlocked_research || [], 
    [pilotStats?.metadata?.unlocked_research]
  );

  useEffect(() => {
    // Update unlocked status based on pilot stats
    setNodes(prev => {
      const hasChanged = prev.some(node => {
        const isUnlocked = unlockedResearch.includes(node.id);
        return node.unlocked !== isUnlocked;
      });
      
      if (!hasChanged) return prev;

      return prev.map(node => ({
        ...node,
        unlocked: unlockedResearch.includes(node.id)
      }));
    });
  }, [unlockedResearch]);

  const handleUnlock = async (nodeId: string) => {
    setLoading(true);
    setError(null);
    try {
      const token = localStorage.getItem('token');
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/game/research/unlock`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
          character_id: pilotStats?.character_id,
          research_id: nodeId
        })
      });

      if (!res.ok) {
        const errData = await res.json();
        throw new Error(errData.error || 'Failed to unlock research');
      }

      const updatedStats = await res.json();
      
      // Update game state
      send({ type: 'UPDATE_PILOT_STATS', stats: updatedStats } as any);
      
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-6 bg-gray-900 text-green-400 font-mono h-full overflow-y-auto border-2 border-green-800 rounded-lg shadow-[0_0_20px_rgba(0,255,0,0.2)]">
      <h2 className="text-2xl font-bold mb-6 border-b border-green-700 pb-2 flex justify-between items-center">
        <span>RESEARCH TERMINAL</span>
        <span className="text-sm">DATA: {research} UNITS</span>
      </h2>

      {error && (
        <div className="mb-4 p-2 bg-red-900/50 border border-red-500 text-red-200 rounded">
          ERROR: {error}
        </div>
      )}

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {nodes.map(node => (
          <div 
            key={node.id}
            className={`p-4 border rounded transition-all duration-300 ${
              node.unlocked 
                ? 'border-green-500 bg-green-900/20 shadow-[0_0_10px_rgba(0,255,0,0.3)]' 
                : 'border-gray-700 bg-gray-800/50 opacity-80 hover:opacity-100'
            }`}
          >
            <div className="flex justify-between items-start mb-2">
              <h3 className="text-lg font-bold">{node.name}</h3>
              {node.unlocked ? (
                <span className="px-2 py-1 bg-green-800 text-green-100 text-xs rounded">UNLOCKED</span>
              ) : (
                <span className="px-2 py-1 bg-gray-700 text-gray-300 text-xs rounded">{node.cost} DATA</span>
              )}
            </div>
            <p className="text-sm text-gray-400 mb-4">{node.description}</p>
            
            {!node.unlocked && (
              <button
                onClick={() => handleUnlock(node.id)}
                disabled={loading || research < node.cost}
                className={`w-full py-2 rounded font-bold text-sm transition-colors ${
                  research >= node.cost
                    ? 'bg-green-700 hover:bg-green-600 text-white shadow-[0_0_10px_rgba(0,255,0,0.5)]'
                    : 'bg-gray-700 text-gray-500 cursor-not-allowed'
                }`}
              >
                {loading ? 'PROCESSING...' : 'UNLOCK'}
              </button>
            )}
          </div>
        ))}
      </div>
    </div>
  );
};
