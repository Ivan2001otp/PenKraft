import React from 'react'
import { CheckCircle2 } from "lucide-react";

export const checklistItems = [
    {
      title: "Dynamic Storylines",
      description:
        "Choose your path and decisions that impact the game world. ",
    },
    {
      title: "Play Anytime, Anywhere",
      description:
        "With cross-platform play, you can team up with friends across different devices—whether they’re on a console, PC expanding your gaming community and playing without limits.",
    },
    {
      title: "Easter Eggs & Secrets",
      description:
        "Games with cleverly hidden secrets or Easter eggs offer an additional layer of fun.",
    },
    {
      title: "In-Game Events",
      description:
        "Limited-time events and seasonal content, like special missions or challenges, bring variety and excitement.",
    },
];

const Workflow = () => {
    return (
        <div className="mt-20">
          <h2 className="text-3xl sm:text-5xl lg:text-6xl text-center mt-6 tracking-wide">
            Accelerate your {" "}
            <span className="bg-gradient-to-r from-orange-500 to-orange-800 text-transparent bg-clip-text">
              Gaming experience & Fun
            </span>
          </h2>
    
          <div className="flex flex-wrap justify-center">
          
            <div className="pt-12 w-full lg:w-1/2">
              {checklistItems.map((item, index) => (
                <div key={index} className="flex mb-12 hover:bg-slate-800 hover:border-white hover:border-2 p-2 my-2 mx-4 rounded-xl hover:scale-90 duration-150 transition">
                  <div className="justify-center items-center rounded-full w-10 h-10 p-2 bg-neutral-900 mx-6 text-green-400">
                    <CheckCircle2 />
                  </div>
                  <div>
                    <h5 className="mt-1 mb-2">{item.title}</h5>
                    <p className="text-md text-neutral-500"> {item.description}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      );
}

export default Workflow