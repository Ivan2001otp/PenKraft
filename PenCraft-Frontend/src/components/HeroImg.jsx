import React from 'react'
import {BsChevronCompactLeft, BsChevronCompactRight} from 'react-icons/bs'
import heroImg from '../assets/cube.jpg';
import {useState} from 'react';
import { RxDotFilled } from 'react-icons/rx';

const HeroImg = () => {
  const slides = [
    {
      url: 'https://images.unsplash.com/photo-1531297484001-80022131f5a1?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2620&q=80',
    },
    {
      url: 'https://images.unsplash.com/photo-1488590528505-98d2b5aba04b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2670&q=80',
    },
    {
      url: 'https://images.unsplash.com/photo-1661961112951-f2bfd1f253ce?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2672&q=80',
    },

    {
      url: 'https://images.unsplash.com/photo-1512756290469-ec264b7fbf87?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2253&q=80',
    },
    {
      url: 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2671&q=80',
    },
  ];

  const [currentIndex, setCurrentIndex] = useState(0);

  const prevSlide = () => {
    const isFirstSlide = (currentIndex === 0);
    const newIndex = (isFirstSlide ? slides.length-1 : currentIndex-1);
    setCurrentIndex(newIndex);
  };


  const nextSlide = () => { 
    const isLastSlide = (currentIndex === slides.length-1);
    const newIndex = (isLastSlide ? 0 : currentIndex + 1);
    setCurrentIndex(newIndex);
  };

  const goToSelectedSlide = (slideIndex) => {
    setCurrentIndex(slideIndex);
  }

  return (
    <div className='max-w-full w-full h-[600px] -mt-12 -z-100 group'>

      <div
        style={{background:`url(${slides[currentIndex].url})`}}
        className='w-full h-full rounded-2xl bg-center bg-cover duration-500'  
      >
      </div>
      
      {/* left arrow */}
      <div  className='
      hidden group-hover:block absolute top-[40%] p-2 translate-x-2 translate-y-[-50%] text-2xl text-white cursor-pointer bg-slate-600 opacity-75 left-5 rounded-full'>
        <BsChevronCompactLeft onClick={prevSlide} size={30} />
      </div>

      {/* right arrow  */}
      <div className='hidden group-hover:block absolute top-[40%] p-2 translate-x-2 translate-y-[-50%] text-2xl text-white cursor-pointer bg-slate-600 opacity-75 right-5 rounded-full'>
        <BsChevronCompactRight size={30} onClick={nextSlide}/>
      </div>

      <div className='flex items-start absolute top-[78%] mx-4  '>
        {
          slides.map((index,item)=>(
            <div
              key={index}
              className='text-2xl cursor-pointer'
              onClick={()=>goToSelectedSlide(index)}
            >
              <RxDotFilled />
            </div>
          ))
        }
      </div>

    </div>
  )
}

export default HeroImg