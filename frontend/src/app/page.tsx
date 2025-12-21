"use client";

import Image from "next/image";
import { useEffect, useState } from "react";
import { LoginButton } from "@/components/login-button";
import { MechCard } from "@/components/mech-card";
import { useAuthSync } from "@/hooks/use-auth-sync";
import { usePrivy } from "@privy-io/react-auth";

export default function Home() {
  const { authenticated, user: privyUser } = usePrivy();
  const { backendToken, user: backendUser, isLoading: isSyncing } = useAuthSync();
  const [mechs, setMechs] = useState<any[]>([]);
  const [isMinting, setIsMinting] = useState(false);

  const fetchMechs = async () => {
    if (backendUser?.id) {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/mechs?user_id=${backendUser.id}`);
      if (res.ok) {
        const data = await res.json();
        setMechs(data || []);
      }
    }
  };

  useEffect(() => {
    if (backendUser?.id) {
      fetchMechs();
    }
  }, [backendUser?.id]);

  const mintStarter = async () => {
    if (!backendUser?.id) return;
    setIsMinting(true);
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/mechs/mint-starter?user_id=${backendUser.id}`, {
        method: 'POST',
      });
      if (res.ok) {
        await fetchMechs();
      }
    } catch (error) {
      console.error("Minting failed:", error);
    } finally {
      setIsMinting(false);
    }
  };

  return (
    <div className="flex min-h-screen flex-col items-center bg-zinc-50 font-sans dark:bg-black text-black dark:text-white">
      <header className="w-full max-w-7xl flex justify-between items-center p-6">
        <div className="flex items-center gap-2">
          <div className="relative h-8 w-8">
            <Image className="dark:invert" src="/next.svg" alt="Logo" fill />
          </div>
          <span className="font-bold tracking-tighter text-xl">PROJECT-0</span>
        </div>
        <LoginButton />
      </header>

      <main className="flex flex-col items-center w-full max-w-7xl px-6 py-12">
        {!authenticated ? (
          <div className="flex flex-col items-center gap-8 text-center mt-20">
            <h1 className="text-5xl font-bold tracking-tighter sm:text-7xl">
              ADAPTIVE UNIVERSE
            </h1>
            <p className="max-w-[600px] text-zinc-600 dark:text-zinc-400 md:text-xl">
              The next generation of AI-driven, Web3-powered space exploration and combat.
            </p>
            <LoginButton />
          </div>
        ) : (
          <div className="w-full space-y-12">
            <section className="flex flex-col md:flex-row justify-between items-start md:items-center gap-6 bg-zinc-900 border border-zinc-800 p-8 rounded-3xl">
              <div>
                <h2 className="text-3xl font-bold mb-2">Welcome, Pilot</h2>
                <p className="text-zinc-400 font-mono text-sm">{privyUser?.wallet?.address}</p>
                <div className="mt-4 flex gap-4">
                  <div className="bg-zinc-800 px-4 py-2 rounded-xl border border-zinc-700">
                    <span className="text-xs text-zinc-500 block uppercase">Credits</span>
                    <span className="font-bold text-indigo-400">{backendUser?.credits || 0} USDT</span>
                  </div>
                </div>
              </div>
              
              {mechs.length === 0 && (
                <button
                  onClick={mintStarter}
                  disabled={isMinting}
                  className="w-full md:w-auto px-8 py-4 bg-indigo-600 hover:bg-indigo-700 disabled:bg-zinc-700 text-white font-bold rounded-2xl shadow-xl shadow-indigo-500/20 transition-all"
                >
                  {isMinting ? 'Assembling Mech...' : 'Claim Starter Mech'}
                </button>
              )}
            </section>

            <section>
              <div className="flex justify-between items-end mb-8">
                <div>
                  <h2 className="text-2xl font-bold">Your Hangar</h2>
                  <p className="text-zinc-500 text-sm">Manage your fleet of Mechs, Tanks, and Ships.</p>
                </div>
                <span className="text-zinc-500 text-sm font-mono">{mechs.length} Units</span>
              </div>

              {mechs.length > 0 ? (
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                  {mechs.map((m) => (
                    <MechCard key={m.id} mech={m} />
                  ))}
                </div>
              ) : (
                <div className="flex flex-col items-center justify-center py-20 border-2 border-dashed border-zinc-800 rounded-3xl bg-zinc-900/30">
                  <p className="text-zinc-500 mb-4">No units found in your hangar.</p>
                  {!isMinting && (
                    <button onClick={mintStarter} className="text-indigo-400 hover:underline font-bold">
                      Claim your first unit to start exploring
                    </button>
                  )}
                </div>
              )}
            </section>
          </div>
        )}
      </main>
      
      <footer className="mt-auto py-10 text-sm text-zinc-500">
        <p>© 2024 Project-0. All rights reserved.</p>
      </footer>
    </div>
  );
}
