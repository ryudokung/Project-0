'use client';

import ShowcaseEngine from '@/components/game/ShowcaseEngine';
import { useAuthSync } from '@/hooks/use-auth-sync';
import { useEffect, useState } from 'react';
import { LoginButton } from '@/components/login-button';
import { usePrivy } from '@privy-io/react-auth';
import { hangarSystem, HangarState } from '@/systems/HangarSystem';
import { gameEvents, GAME_EVENTS } from '@/systems/EventBus';

interface HangarProps {
  onDeploy: () => void;
  onGacha: () => void;
}

export default function Hangar({ onDeploy, onGacha }: HangarProps) {
  const { user: backendUser } = useAuthSync();
  const { linkWallet } = usePrivy();
  const [hangarState, setHangarState] = useState<HangarState>(hangarSystem.getState());
  const [isLinking, setIsLinking] = useState(false);

  useEffect(() => {
    const unsubscribe = gameEvents.on(GAME_EVENTS.HANGAR_UPDATED, (newState: HangarState) => {
      setHangarState(newState);
    });

    if (backendUser?.active_character_id || backendUser?.id) {
      hangarSystem.refreshMechs(backendUser.active_character_id || backendUser.id);
    }

    return () => unsubscribe();
  }, [backendUser?.id, backendUser?.active_character_id]);

  const handleLinkWallet = async () => {
    setIsLinking(true);
    try {
      await linkWallet();
    } catch (error) {
      console.error("Failed to link wallet:", error);
    } finally {
      setIsLinking(false);
    }
  };

  const currentMech = hangarState.mechs.find(m => m.id === hangarState.selectedMechId) || hangarState.mechs[0];

  return (
    <div className="relative w-full h-screen bg-black">
      <ShowcaseEngine mechId={currentMech?.id || "MCH-001-ALPHA"} />
      
      {/* UI Overlay */}
      <div className="absolute top-8 left-8 z-10">
        <div className="flex items-start gap-4">
          {/* Pilot ID Badge */}
          <div className="bg-zinc-900/80 border border-zinc-700 p-4 rounded-lg backdrop-blur-md shadow-xl flex gap-4 items-center">
            <div className="w-16 h-16 bg-zinc-800 border border-zinc-700 rounded overflow-hidden flex items-center justify-center">
              {backendUser?.active_character?.gender ? (
                <img 
                  src={`/images/pilots/${backendUser.active_character.gender.toLowerCase()}.png`} 
                  alt="Pilot Avatar"
                  className="w-full h-full object-cover"
                />
              ) : (
                <div className="text-pink-500 text-2xl font-bold">
                  {backendUser?.active_character?.name?.[0] || 'P'}
                </div>
              )}
            </div>
            <div>
              <h2 className="text-xl font-bold text-white tracking-tighter uppercase italic">
                {backendUser?.active_character?.name || 'UNREGISTERED PILOT'}
              </h2>
              <div className="flex gap-2 mt-1">
                <span className="text-[10px] font-mono text-zinc-400 bg-zinc-800 px-1.5 py-0.5 rounded">
                  {backendUser?.active_character?.gender || 'UNKNOWN'}
                </span>
                <span className="text-[10px] font-mono text-pink-400 bg-pink-400/10 border border-pink-400/20 px-1.5 py-0.5 rounded">
                  RANK 1
                </span>
              </div>
            </div>
          </div>
        </div>
        
        <div className="mt-4 flex flex-col gap-2">
          <LoginButton />
          
          {backendUser && !backendUser.wallet_address && (
            <button 
              onClick={handleLinkWallet}
              disabled={isLinking}
              className="bg-indigo-600 text-white px-4 py-2 text-xs font-bold uppercase tracking-widest hover:bg-indigo-500 transition-all disabled:opacity-50"
            >
              {isLinking ? 'Linking...' : 'Link External Wallet'}
            </button>
          )}

          {backendUser?.wallet_address && (
            <div className="text-[10px] text-green-500 font-mono bg-green-500/10 border border-green-500/20 px-2 py-1 rounded">
              WALLET: {backendUser.wallet_address.slice(0, 6)}...{backendUser.wallet_address.slice(-4)}
            </div>
          )}
        </div>
        {currentMech && (
          <div className="mt-4 flex gap-2">
            {currentMech.is_void_touched && (
              <span key="void-touched" className="px-2 py-1 bg-purple-900/50 border border-purple-500 text-purple-300 text-[10px] font-bold uppercase tracking-widest">
                Void-Touched
              </span>
            )}
            <span key="tier" className="px-2 py-1 bg-zinc-900 border border-zinc-700 text-zinc-400 text-[10px] font-bold uppercase tracking-widest">
              Tier {currentMech.tier || 1}
            </span>
            <span key="rarity" className="px-2 py-1 bg-zinc-900 border border-zinc-700 text-zinc-400 text-[10px] font-bold uppercase tracking-widest">
              {currentMech.rarity}
            </span>
          </div>
        )}
      </div>

      <div className="absolute bottom-8 right-8 z-10 flex flex-col gap-4">
        <button 
          onClick={onGacha}
          className="bg-yellow-400 text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-yellow-300 transition-colors text-center"
        >
          Void Signals (Gacha)
        </button>
        <button className="bg-white text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-gray-200 transition-colors">
          Share to Social
        </button>
        <button 
          onClick={onDeploy}
          className="border border-white text-white px-6 py-2 font-bold uppercase tracking-tighter hover:bg-white hover:text-black transition-colors text-center"
        >
          Launch Expedition
        </button>
      </div>
    </div>
  );
}
