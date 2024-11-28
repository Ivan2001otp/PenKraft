import React from "react";
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
    const isFirstSlide = currentIndex === 0;
    const newIndex = isFirstSlide ? slides.length - 1 : currentIndex - 1;
    setCurrentIndex(newIndex);
  };

  const nextSlide = () => {
    const isLastSlide = currentIndex === slides.length - 1;
    const newIndex = isLastSlide ? 0 : currentIndex + 1;
    setCurrentIndex(newIndex);
  };

  const goToSelectedSlide = (slideIndex) => {
    console.log(slideIndex);
    setCurrentIndex(slideIndex);
  };

  return (
    <div className="max-w-[1540px] h-[800px] m-auto md:-mt-14  -z-100 group">
      {/* <div
        style={{ background: `url(${slides[currentIndex].url})` }}
        className="w-full h-full rounded-2xl bg-center bg-contain duration-500"
      ></div> */}
      <img
        src={slides[currentIndex].url}
        className="flex w-full md:h-[720px] duration-500"
      />

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

      <div className="hidden md:flex items-start absolute bottom-[5%] mx-4  ">
        {slides.map((index, item) => (
          <div
            key={index}
            className="text-2xl cursor-pointer"
            onClick={() => goToSelectedSlide(index)}
          >
            <RxDotFilled />
          </div>
        ))}
      </div>
    </div>
  );
};

export default HeroImg;
