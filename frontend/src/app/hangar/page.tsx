import ShowcaseEngine from '@/components/game/ShowcaseEngine';

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
        </div>

        <div className="absolute bottom-8 right-8 z-10 flex flex-col gap-4">
          <button className="bg-white text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-gray-200 transition-colors">
            Share to Social
          </button>
          <button className="border border-white text-white px-6 py-2 font-bold uppercase tracking-tighter hover:bg-white hover:text-black transition-colors">
            Enter Cockpit
          </button>
        </div>
      </div>
    </main>
  );
}
