'use client';

import ShowcaseEngine from '@/components/game/ShowcaseEngine';
import Link from 'next/link';
import { useAuthSync } from '@/hooks/use-auth-sync';
import { useEffect, useState } from 'react';
import { LoginButton } from '@/components/login-button';
import { usePrivy } from '@privy-io/react-auth';

export default function HangarPage() {
  const { user: backendUser } = useAuthSync();
  const { linkWallet, user: privyUser } = usePrivy();
  const [mechs, setMechs] = useState<any[]>([]);
  const [isLinking, setIsLinking] = useState(false);

  useEffect(() => {
    if (backendUser?.id) {
      fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/mechs?user_id=${backendUser.id}`)
        .then(res => res.json())
        .then(data => setMechs(data || []));
    }
  }, [backendUser?.id]);

  const handleLinkWallet = async () => {
    setIsLinking(true);
    try {
      await linkWallet();
      // After linking, we need to sync with backend
      // Privy will update privyUser, which triggers useAuthSync
    } catch (error) {
      console.error("Failed to link wallet:", error);
    } finally {
      setIsLinking(false);
    }
  };

  const currentMech = mechs[0];

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-0 bg-black">
      <div className="relative w-full h-screen">
        <ShowcaseEngine mechId={currentMech?.id || "MCH-001-ALPHA"} />
        
        {/* UI Overlay for Flexing */}
        <div className="absolute top-8 left-8 z-10">
          <h1 className="text-4xl font-bold text-white tracking-tighter uppercase italic">
            The Hangar
          </h1>
          <p className="text-gray-400 font-mono text-sm mt-2">
            [STATUS: SECURE] // [OWNER: {backendUser?.username || 'PILOT_0'}]
          </p>
          
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
          <Link href="/gacha" className="bg-yellow-400 text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-yellow-300 transition-colors text-center">
            Void Signals (Gacha)
          </Link>
          <button className="bg-white text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-gray-200 transition-colors">
            Share to Social
          </button>
          <Link href="/exploration-loop" className="border border-white text-white px-6 py-2 font-bold uppercase tracking-tighter hover:bg-white hover:text-black transition-colors text-center">
            Launch Expedition
          </Link>
        </div>
      </div>
    </main>
  );
}
