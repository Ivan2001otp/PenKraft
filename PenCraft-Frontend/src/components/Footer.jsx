import React from 'react'
import penkraftLogo from "../assets/penkraft.png"


const resourcesLinks = [
    { href: "#", text: "Getting Started" },
    { href: "#", text: "Documentation" },
    { href: "#", text: "Tutorials" },
    { href: "#", text: "API Reference" },
    { href: "#", text: "Community Forums" },
  ];
  
 const platformLinks = [
    { href: "#", text: "Features" },
    { href: "#", text: "Supported Devices" },
    { href: "#", text: "System Requirements" },
    { href: "#", text: "Downloads" },
    { href: "#", text: "Release Notes" },
  ];
  
const communityLinks = [
    { href: "#", text: "Events" },
    { href: "#", text: "Meetups" },
    { href: "#", text: "Conferences" },
    { href: "#", text: "Hackathons" },
    { href: "#", text: "Jobs" },
  ];

const Footer = () => {

  return (
    <footer className='mt-20 mb-20 border-t border-neutral-700'>
        <div className='grid grid-cols-2 lg:grid-cols-4 gap-2 ml-[10rem]'>
            <div className='mt-20'>
                <h3 className='underline-effect text-md font-semibold mb-4'>
                    Resources
                </h3>

                <ul className='space-y-2'>
                    {
                        resourcesLinks.map((resource,index) => (
                                <li key={index}>
                                    <a className='text-neutral-300 hover:text-white' href={resource.href}>{resource.text}</a>
                                </li>
                        ))
                    }
                </ul>
            </div>

            <div className='mt-20'>
                <h3 className='underline-effect text-md font-semibold mb-4'>Platform</h3>
                <ul className='space-y-2'>
                    {
                        platformLinks.map((link,index) => (
                            <li
                                key={index}
                            >
                                <a href={link.href}
                                    className='text-neutral-300 hover:text-white'
                                >{link.text}</a>
                            </li>
                        ))
                    }
                </ul>
            </div>

            <div className='mt-20'>
                <h3 className='underline-effect text-md font-semibold mb-4'>
                            Community
                </h3>
                <ul className='space-y-2'>
                    {
                        communityLinks.map((item,index) => (
                            <li
                                key={index}
                            >
                                <a className='text-neutral-300hover:text-white'>
                                    {item.text}
                                </a>
                            </li>
                        ))
                    }
                </ul>
            </div>

            <div className='mt-20'>
                <img className='h-40 w-40 rounded-lg' src={penkraftLogo}/>
            </div>

        </div>
    </footer>
  )
}

export default Footer