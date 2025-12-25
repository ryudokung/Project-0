'use client';

import { Canvas } from '@react-three/fiber';
import { OrbitControls, PerspectiveCamera, Environment, ContactShadows, Float } from '@react-three/drei';
import { Suspense } from 'react';

interface ShowcaseEngineProps {
  vehicleId?: string;
  visualDna?: string;
}

export default function ShowcaseEngine({ vehicleId, visualDna }: ShowcaseEngineProps) {
  return (
    <div className="w-full h-full bg-black">
      <Canvas shadows dpr={[1, 2]}>
        <PerspectiveCamera makeDefault position={[0, 2, 5]} fov={50} />
        <OrbitControls 
          enablePan={false} 
          minDistance={3} 
          maxDistance={10} 
          maxPolarAngle={Math.PI / 2} 
        />
        
        <Suspense fallback={null}>
          <Environment preset="city" />
          
          <Float speed={1.5} rotationIntensity={0.5} floatIntensity={0.5}>
            {/* Placeholder for the Vehicle Model */}
            <mesh castShadow receiveShadow position={[0, 1, 0]}>
              <boxGeometry args={[1, 2, 1]} />
              <meshStandardMaterial color="#333" metalness={0.8} roughness={0.2} />
            </mesh>
          </Float>

          <ContactShadows 
            position={[0, 0, 0]} 
            opacity={0.4} 
            scale={10} 
            blur={2} 
            far={4.5} 
          />
          
          <gridHelper args={[20, 20, '#222', '#111']} />
        </Suspense>

        <ambientLight intensity={0.5} />
        <spotLight position={[10, 10, 10]} angle={0.15} penumbra={1} intensity={1} castShadow />
      </Canvas>
      
      <div className="absolute bottom-4 left-4 text-white font-mono text-xs uppercase tracking-widest">
        Showcase Engine v1.0 // Vehicle: {vehicleId || 'Unknown'}
      </div>
    </div>
  );
}
