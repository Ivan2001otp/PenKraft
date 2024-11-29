import React, { useEffect } from "react";
import { BsChevronCompactLeft, BsChevronCompactRight } from "react-icons/bs";
import heroImg from "../assets/cube.jpg";
import { useState } from "react";
import { RxDotFilled } from "react-icons/rx";

import img1 from "../assets/assassins.jpeg";
import img2 from "../assets/avatar.jpeg";
import img3 from "../assets/chronicles.jpeg";
import img4 from "../assets/magicimg.jpg";
import img5 from "../assets/ragnarok.jpg";
import img6 from "../assets/treeimg.jpg";
import img14 from "../assets/sekiero.jpg";
import img8 from "../assets/blackops.jpg";
import img9 from "../assets/spidy1.jpg";
import img10 from "../assets/spidy2.jpg";
import img11 from "../assets/ww3.jpg";
import img12 from "../assets/cod.jpg";
import img13 from "../assets/wukong.jpg";

const HeroImg = () => {
  const slides = [
    {
      url: img1,
    },
    {
      url: img2,
    },
    {
      url: img3,
    },

    {
      url: img4,
    },
    {
      url: img5,
    },
    {
      url: img6,
    },
    {
      url: img14,
    },
    {
      url: img8,
    },
    {
      url: img9,
    },
    {
      url: img10,
    },
    {
      url: img11,
    },
    {
      url: img12,
    },
    {
      url: img13,
    },
  ];

  const [currentIndex, setCurrentIndex] = useState(0);

  const prevSlide = () => {
    // const isFirstSlide = currentIndex === 0;
    // const newIndex = isFirstSlide ? slides.length - 1 : currentIndex - 1;

    setCurrentIndex(
      (prevIndex) => (prevIndex - 1 + slides.length) % slides.length
    );
  };

  const nextSlide = () => {
    // const isLastSlide = currentIndex === slides.length - 1;
    // const newIndex = isLastSlide ? 0 : currentIndex + 1;
    setCurrentIndex((prevIndex) => (prevIndex + 1) % slides.length);
  };

  useEffect(() => {
    const interval = setInterval(() => {
      nextSlide();
    }, 5000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div className="flex flex-col mt-8 md:mt-0 lg:mt-0">
      <div
        style={{
          backgroundImage: `url(${slides[currentIndex].url})`,
          marginTop: `-8rem`,
        }}
        className="bg-center bg-cover duration-500 w-full h-[16rem] md:h-[760px]"
      >
        <h1 className="hidden md:block lg:ml-8 w-fit bokor-regular tracking-wide opacity-60 hover:opacity-100 duration-300  text-3xl md:text-5xl lg:text-6xl lg:mt-[38%]">
          From Gaming to{" "}
          <span className="bokor-regular hover:text-8xl transistion-all duration-300 ease-in-out text-transparent bg-gradient-to-r from-orange-500 to-red-800 bg-clip-text">
            Blogging
          </span>
        </h1>

        <h1 className="hidden md:block lg:ml-8 w-fit bokor-regular text-3xl sm:text-5xl lg:text-6xl text-left tracking-wide opacity-60 hover:opacity-100 duration-300">
          your{" "}
          <span className="bokor-regular hover:text-7xl transistion-all duration-300 ease-in-out text-transparent bg-gradient-to-r from-yellow-500 to-orange-600 bg-clip-text  hover:underline hover:underline-offset-4 hover:decoration-slate-500">
            Game
          </span>{" "}
          ,your{" "}
          <span className="bokor-regular hover:text-7xl transistion-all duration-300 ease-in-out text-transparent bg-gradient-to-r from-red-500 to-orange-400 bg-clip-text  hover:underline hover:underline-offset-4 hover:decoration-slate-500">
            Story
          </span>
        </h1>

        <div className="w-full bg-gradient-to-t from-neutral-900 via-neutral-800 to-transparent"></div>
      </div>

      <div className="mt-4 mx-auto md:hidden">
        <h1 className="bokor-regular text-3xl text-left tracking-wide opacity-60 ">
          From Gaming to{" "}
          <span className="bokor-regular bg-gradient-to-r from-orange-500 to-red-800 bg-clip-text">
            Blogging
          </span>
        </h1>

        <h1 className="bokor-regular text-3xl text-left tracking-wide opacity-60 ">
          your{" "}
          <span className="bokor-regular hover:text-7xl text-transparent bg-gradient-to-r from-yellow-500 to-orange-600 bg-clip-text">
            Game
          </span>{" "}
          ,your{" "}
          <span className="bokor-regular hover:text-7xl transistion-all duration-300 ease-in-out text-transparent bg-gradient-to-r from-red-500 to-orange-400 bg-clip-text  hover:underline hover:underline-offset-4 hover:decoration-slate-500">
            Story
          </span>
        </h1>
      </div>

      {/* <div className="flex justify-center md:-mt-[8rem]">
        <img className="bg-center bg-cover duration-500 justify-center "  src={slides[currentIndex].url}/>
      </div> */}

      {/* <div className="hidden md:block translate-y-[-100%] w-full h-[300px] bg-gradient-to-t from-neutral-950 to-transparent border-b-2"></div> */}

      {/* <div className="translate-x-8  absolute mt-2 flex flex-col items-center text-center lg:bottom-[10%]">
        <h1 className="mb-4 bokor-regular text-3xl sm:text-5xl lg:text-6xl text-left tracking-wide opacity-60 hover:opacity-100 duration-300">
          From Gaming to{" "}
          <span className="bokor-regular hover:text-8xl transistion-all duration-300 ease-in-out text-transparent bg-gradient-to-r from-blue-700 to-red-800 bg-clip-text hover:underline hover:underline-offset-4 hover:decoration-indigo-200">
            Blogging
          </span>
        </h1>
        <h1 className="bokor-regular text-3xl sm:text-5xl lg:text-6xl text-left tracking-wide opacity-60 hover:opacity-100 duration-300">
          your{" "}
          <span className="bokor-regular hover:text-7xl transistion-all duration-300 ease-in-out text-transparent bg-gradient-to-r from-yellow-500 to-orange-600 bg-clip-text  hover:underline hover:underline-offset-4 hover:decoration-slate-500">
            Game
          </span>{" "}
          ,your{" "}
          <span className="bokor-regular hover:text-7xl transistion-all duration-300 ease-in-out text-transparent bg-gradient-to-r from-red-500 to-orange-400 bg-clip-text  hover:underline hover:underline-offset-4 hover:decoration-slate-500">
            Story
          </span>
        </h1>
      </div> */}
    </div>
  );
};

export default HeroImg;
