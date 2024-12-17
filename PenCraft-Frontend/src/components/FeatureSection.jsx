import React from 'react'
import { BotMessageSquare } from "lucide-react";
import { BatteryCharging } from "lucide-react";
import { Fingerprint } from "lucide-react";
import { ShieldHalf } from "lucide-react";
import { PlugZap } from "lucide-react";
import { GlobeLock } from "lucide-react";

const features = [
    {
        id:1,
        text: "In Depth Game Guides",
        description : "Our platform provides detailed, well-written guides for RPG and FPS games, covering everything from the main story to side quests and hidden secrets from dope gamers. These guides are perfect for gamers who want to progress through challenging levels, discover new strategies, or unlock the full potential of their games. Whether you're a newbie or a seasoned pro, these guides will help you level up your skills and enjoy the game to the fullest"
    },
    {
        id:2,
        text: "Boss Fight Solutions",
        description : "Struggling with a tough boss fight? Our site offers comprehensive boss fight strategies, tailored for both RPG and FPS games. Get step-by-step instructions on how to defeat the hardest bosses, including recommended weapons, techniques, and tips to overcome their unique challenges. With our detailed solutions, you'll defeat even the most intimidating foes with ease."
    },
    {
        id:3,
        text: "Game Reviews",
        description : "Get honest and thorough game reviews from passionate players who know RPG and FPS games inside and out. Our reviews cover everything from gameplay mechanics to story depth, graphics, and multiplayer modes, helping you decide whether a game is worth your time. Stay up-to-date with the latest releases and find the perfect game for your next adventure."
    },
    {
        id:4,
        text: "Latest News & Updates",
        description : "Stay informed about the latest game releases, patches, updates, and news in the RPG and FPS worlds. Our platform provides real-time coverage on the most important gaming events, ensuring you're always in the loop. Whether it’s a new game launch or a major patch, you'll find the information you need to stay ahead in the gaming world."
    },
    {
        id:5,
        text: "New Levels, New Challenges",
        description : "Stay informed about the latest game releases, patches, updates, and news in the RPG and FPS worlds. Our platform provides real-time coverage on the most important gaming events, ensuring you're always in the loop. Whether it’s a new game launch or a major patch, you'll find the information you need to stay ahead in the gaming world."
    },
    {
        id:6,
        text: "Game Hacks & Secrets",
        description : "Stay informed about the latest game releases, patches, updates, and news in the RPG and FPS worlds. Our platform provides real-time coverage on the most important gaming events, ensuring you're always in the loop. Whether it’s a new game launch or a major patch, you'll find the information you need to stay ahead in the gaming world."
    }
    
]


function getIconById(Id) {
    if (Id === 1) {
      return <BotMessageSquare />;
    } else if (Id === 2) {
      return <Fingerprint />;
    } else if (Id === 3) {
      return <ShieldHalf />;
    } else if (Id === 4) {
      return <BatteryCharging />;
    } else if (Id === 5) {
      return <PlugZap />;
    } else if (Id === 6) {
      return <GlobeLock />;
    }
    return "0";
  }

const FeatureSection = () => {
  return (
    <div className='mt-20  border-neutral-800 min-h-[800px]'>
        <div className='text-center mt-6'>
            <span className='uppercase text-orange-500 rounded-full h-10 font-medium sm:text-3xl lg:text-5xl text-xl px-2 py-1 logo-font'>Features</span>
            {/* Conquer Boss Fights, Clear Levels, and Share Your Experience! */}
            <h2 className='text-3xl sm:text-5xl lg:text-6xl lg:mt-20 tracking-wide bokor-regular'>
                Conquer Boss Fights, Clear Levels and {' '}
                <span className='bg-gradient-to-r from-orange-500 to-orange-800 text-transparent bg-clip-text bokor-regular'>Share your Experience</span>
            </h2>


            <div className="flex flex-wrap mt-10 lg:mt-16 p-2">
                {
                    features.map((item, index) => (
                        <div key={index} className='w-full hover:border-yellow-400 sm:w-1/2 lg:w-1/3 bg-slate-900 my-1 rounded-sm border-2 border-yellow-700'>
                            <div className='flex mt-2 ml-2'>
                                {getIconById(item.id)}
                            </div>

                            <div>
                                <h5 className='mt-1text-xl text-orange-600 text-left p-2 logo-font'>{item.text}</h5>
                                <p className='text-left p-2 mb-4 logo-font'>{item.description}</p>
                            </div>
                        </div>
                    ))
                }
            </div>
        </div>
    </div>
  )
}

export default FeatureSection