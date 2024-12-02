import React from 'react'
import bg_vid from "../assets/bg_video.mp4"

const HeroSection = () => {
  return (
    <div className='flex justify-center mt-10 '>
      <video
        autoPlay
        loop
        muted
        className='rounded-lg border-white border-b-2 border-t-2 mx-4 my-4'
      >
        <source src={bg_vid} type='video/mp4'/>
        Your browser does not support the video tag.

        
      </video>
    </div>
  )
}

export default HeroSection