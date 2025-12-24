import ShowcaseEngine from '@/components/game/ShowcaseEngine';
import Link from 'next/link';

export default function HangarPage() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-0 bg-black">
      <div className="relative w-full h-screen">
        <ShowcaseEngine mechId="MCH-001-ALPHA" />
        
        {/* UI Overlay for Flexing */}
        <div className="absolute top-8 left-8 z-10">
          <h1 className="text-4xl font-bold text-white tracking-tighter uppercase italic">
            The Hangar
          </h1>
          <p className="text-gray-400 font-mono text-sm mt-2">
            [STATUS: SECURE] // [OWNER: PILOT_0]
          </p>
          <div className="mt-4 flex gap-2">
            <span className="px-2 py-1 bg-purple-900/50 border border-purple-500 text-purple-300 text-[10px] font-bold uppercase tracking-widest">
              Void-Touched
            </span>
            <span className="px-2 py-1 bg-zinc-900 border border-zinc-700 text-zinc-400 text-[10px] font-bold uppercase tracking-widest">
              Tier 2
            </span>
          </div>
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
