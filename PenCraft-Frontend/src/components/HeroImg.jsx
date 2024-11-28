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
import img7 from "../assets/cube.jpg";

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
      url: img7,
    },
  ];

  const [currentIndex, setCurrentIndex] = useState(0);

  const prevSlide = () => {
    // const isFirstSlide = currentIndex === 0;
    // const newIndex = isFirstSlide ? slides.length - 1 : currentIndex - 1;

    setCurrentIndex((prevIndex)=>(prevIndex-1 + slides.length)%slides.length);
  };

  const nextSlide = () => {
    // const isLastSlide = currentIndex === slides.length - 1;
    // const newIndex = isLastSlide ? 0 : currentIndex + 1;
    setCurrentIndex((prevIndex) => (prevIndex+1) % slides.length);
  };

  const goToSelectedSlide = (slideIndex) => {
    console.log(slideIndex);
    setCurrentIndex(slideIndex);
  };

  useEffect(()=>{
    const interval = setInterval(()=>{
      nextSlide();
    },3000);

    return ()=>clearInterval(interval);
  },[]);

  return (
    <div className="max-w-[1540px] h-[790px] m-auto md:-mt-14  -z-100 group ">
        {/* <div 
         
        className="transition-transform duration-500 flex">
          <img
            key={currentIndex}
            src={slides[currentIndex].url}
            className="flex w-full md:h-fit"
          />
        </div> */}
        <div
        style={{ backgroundImage: `url(${slides[currentIndex].url})` }}
        className='max-w-full h-[180px] md:h-full flex bg-center bg-cover md:bg-contain duration-500'
      ></div>

      {/* left arrow */}
      <div
        className="
      hidden group-hover:block absolute top-[20%]  md:top-[50%] p-2 -translate-x-1 translate-y-[-50%] text-xl md:text-2xl text-white cursor-pointer bg-slate-600 opacity-75 left-4 rounded-full"
      >
        <BsChevronCompactLeft onClick={prevSlide} size={20} />
      </div>

      {/* right arrow  */}
      <div className="hidden group-hover:block absolute top-[20%] md:top-[50%] p-2 translate-x-1 translate-y-[-50%] text-xl md:text-2xl text-white cursor-pointer bg-slate-600 opacity-75 right-5 rounded-full">
        <BsChevronCompactRight size={20} onClick={nextSlide} />
      </div>

      <div className="hidden md:flex items-start absolute  lg:bottom-[-5%] mx-4  ">
        {slides.map((index, _) => (
          <div
            key={index}
            className="text-xl cursor-pointer"
            onClick={() => goToSelectedSlide(index)}
          >
            <RxDotFilled/>
          </div>
        ))}
      </div>

      <div className="translate-x-8 absolute mt-2 flex flex-col items-center text-center lg:bottom-[10%]">
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
      </div>
    </div>
  );
};

export default HeroImg;
