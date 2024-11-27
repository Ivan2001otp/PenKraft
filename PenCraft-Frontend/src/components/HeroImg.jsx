import React from 'react'
import heroImg from '../assets/heroimage.jpg';

const HeroImg = () => {
  return (
    <div className='flex flex-wrap -mt-16'>
        <img
            src={heroImg}
            alt='HeroImage'
            className='h-1/3 w-full'
        />
    </div>
  )
}

export default HeroImg