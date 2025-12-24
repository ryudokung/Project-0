'use client';

import { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { useAuthSync } from '@/hooks/use-auth-sync';
import { LoginButton } from '@/components/login-button';

interface GachaResult {
  item_type: string;
  rarity: string;
  item: any;
}

export default function GachaPage() {
  const { user } = useAuthSync();
  const [isPulling, setIsPulling] = useState(false);
  const [results, setResults] = useState<GachaResult[] | null>(null);
  const [showReveal, setShowReveal] = useState(false);

  const handlePull = async (count: number, type: string = 'VOID_SIGNAL') => {
    if (!user?.id) {
      alert('Please login first');
      return;
    }
    setIsPulling(true);
    setResults(null);
    
    try {
      const res = await fetch('http://localhost:8080/api/v1/gacha/pull', { 
        method: 'POST', 
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ 
          count, 
          pull_type: type,
          user_id: user.id
        }) 
      });
      
      if (!res.ok) {
        const errorData = await res.json();
        throw new Error(errorData.error || 'Failed to pull');
      }

      const data = await res.json();
      setResults(data.results);
      setIsPulling(false);
      setShowReveal(true);
    } catch (error: any) {
      alert(error.message);
      setIsPulling(false);
    }
  };

  const rollMockRarity = () => {
    const r = Math.random() * 100;
    if (r < 1) return 'SINGULARITY';
    if (r < 5) return 'RELIC';
    if (r < 15) return 'PROTOTYPE';
    if (r < 40) return 'REFINED';
    return 'STANDARD';
  };

  const getRarityColor = (rarity: string) => {
    switch (rarity) {
      case 'SINGULARITY': return 'text-yellow-400 shadow-[0_0_20px_rgba(251,191,36,0.5)]';
      case 'RELIC': return 'text-purple-500';
      case 'PROTOTYPE': return 'text-blue-400';
      case 'REFINED': return 'text-green-400';
      default: return 'text-white';
    }
  };

  const isHighRarity = results?.some(r => r.rarity === 'RELIC' || r.rarity === 'SINGULARITY');

  return (
    <main className={`min-h-screen bg-black text-white flex flex-col items-center justify-center p-8 overflow-hidden ${isHighRarity && showReveal ? 'animate-glitch' : ''}`}>
      <div className="absolute top-8 right-8 z-20">
        <LoginButton />
      </div>
      <style>{`
        @keyframes glitch {
          0% { transform: translate(0) }
          20% { transform: translate(-2px, 2px) }
          40% { transform: translate(-2px, -2px) }
          60% { transform: translate(2px, 2px) }
          80% { transform: translate(2px, -2px) }
          100% { transform: translate(0) }
        }
        .animate-glitch {
          animation: glitch 0.2s infinite;
          filter: hue-rotate(90deg);
        }
      `}</style>
      <div className="absolute inset-0 bg-[radial-gradient(circle_at_center,_var(--tw-gradient-stops))] from-zinc-900 via-black to-black opacity-50" />
      
      <h1 className="text-6xl font-black tracking-tighter italic uppercase mb-12 z-10">
        Void Signals
      </h1>

      <div className="flex gap-8 z-10">
        <button 
          onClick={() => handlePull(1, 'DAILY_SIGNAL')}
          disabled={isPulling}
          className="group relative px-8 py-4 bg-zinc-900 border border-zinc-800 hover:border-yellow-400 transition-all"
        >
          <span className="relative z-10 font-bold uppercase tracking-widest text-yellow-400">Daily Signal</span>
          <div className="absolute inset-0 bg-yellow-400 opacity-0 group-hover:opacity-5 transition-opacity" />
        </button>

        <button 
          onClick={() => handlePull(1)}
          disabled={isPulling}
          className="group relative px-12 py-4 bg-zinc-900 border border-zinc-800 hover:border-white transition-all"
        >
          <span className="relative z-10 font-bold uppercase tracking-widest">Decode x1</span>
          <div className="absolute inset-0 bg-white opacity-0 group-hover:opacity-5 transition-opacity" />
        </button>

        <button 
          onClick={() => handlePull(10)}
          disabled={isPulling}
          className="group relative px-12 py-4 bg-white text-black border border-white hover:bg-zinc-200 transition-all"
        >
          <span className="relative z-10 font-bold uppercase tracking-widest">Decode x10</span>
        </button>
      </div>

      <AnimatePresence>
        {isPulling && (
          <motion.div 
            key="pulling-overlay"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            className="fixed inset-0 z-50 bg-black flex flex-col items-center justify-center"
          >
            <motion.div 
              animate={{ 
                scale: [1, 1.2, 1],
                rotate: [0, 180, 360],
                opacity: [0.5, 1, 0.5]
              }}
              transition={{ duration: 2, repeat: Infinity }}
              className="w-32 h-32 border-2 border-white rounded-full flex items-center justify-center"
            >
              <div className="w-16 h-16 bg-white rounded-full blur-xl" />
            </motion.div>
            <p className="mt-8 font-mono text-sm tracking-[0.5em] uppercase animate-pulse">
              Scanning Void Frequencies...
            </p>
          </motion.div>
        )}

        {showReveal && results && (
          <motion.div 
            key="reveal-overlay"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            className="fixed inset-0 z-50 bg-black/95 backdrop-blur-xl flex flex-col items-center justify-center p-12"
          >
            <div className="grid grid-cols-5 gap-4 max-w-6xl w-full">
              {results.map((res, i) => (
                <motion.div
                  key={`result-${i}-${res.item_type}`}
                  initial={{ y: 50, opacity: 0 }}
                  animate={{ y: 0, opacity: 1 }}
                  transition={{ delay: i * 0.1 }}
                  className="aspect-[2/3] bg-zinc-900 border border-zinc-800 p-4 flex flex-col items-center justify-between group hover:border-white transition-colors"
                >
                  <div className={`text-[10px] font-mono uppercase tracking-widest ${getRarityColor(res.rarity)}`}>
                    {res.rarity}
                  </div>
                  <div className="w-full h-32 bg-zinc-800 rounded-sm overflow-hidden relative">
                    <div className="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent" />
                  </div>
                  <div className="text-center">
                    <div className="text-xs font-bold uppercase">{res.item_type}</div>
                    <div className="text-[10px] text-zinc-500 font-mono mt-1">ID: INTERCEPTED_{i}</div>
                  </div>
                </motion.div>
              ))}
            </div>

            <button 
              onClick={() => setShowReveal(false)}
              className="mt-16 px-12 py-3 border border-white text-white font-bold uppercase tracking-widest hover:bg-white hover:text-black transition-all"
            >
              Confirm
            </button>
          </motion.div>
        )}
      </AnimatePresence>
    </main>
  );
}
