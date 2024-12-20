import React, { useState } from "react";
import { Menu, X } from "lucide-react";
import penkraft from "../assets/PenkraftLogov2.png";
import twitter from "../assets/twitter.png";

const navItems = [
  { label: "Home", href: "#" },
  { label: "About", href: "#" },
  { label: "Contact", href: "#" },
  { label: "Articles", href: "#" },
  { label: "Smash", href: "#" },
  { label: "Social-Media", href: twitter },
];

const Navbar = () => {
  const [mobileOpenDrawer, setMobileOpenDrawer] = useState(false);

  const toggleNavbar = () => {
    setMobileOpenDrawer(!mobileOpenDrawer);
  };

  return (
    <nav className="sticky top-0 z-80 py-3 backdrop-blur-lg bg-neutral-950 bg-opacity-15 border-neutral-700/80 rounded-b-xl">
      <div className="container relative px-4 mx-auto md:text-sm">
        <div className="flex justify-between items-center">
          {/* logo  */}
          <div className="flex flex-shrink-0 items-center">
            <img
              className="h-12 w-12 rounded-md mr-3"
              src={penkraft}
              alt="logo"
            />
            <span style={{"--underline-color":'#f3815f'}} className="logo-font text-2xl tracking-light logo-underline">PenKraft</span>
          </div>

          <ul className="hidden lg:flex space-x-10 items-center">
            {navItems.map((item, index) => (
              <li
                style={{"--underline-color":"#f3efec"}}
                className="logo-font logo-underline  hover:text-yellow-500 transition duration-400"
                key={index}
              >
                {item.label === "Social-Media" ? (
                  <a href="https://x.com/I_Dsouza7697" >
                  <img
                    className="h-8 w-8 cursor-pointer bg-white 
                                hover:bg-orange-400 transition duration-300
                                rounded-md"
                    src={twitter}
                    alt="Creator Profile"
                    
                  />
                  </a>
                  
                ) : (
                  <a className="text-[16px]" href={item.href}>
                    {item.label}
                  </a>
                )}
              </li>
            ))}
          </ul>

          <div className="hidden lg:flex justify-center items-center space-x-10">
            <a href="#" className="py-2 px-3 border-b-2  rounded-md">
              Log In
            </a>
            <a
              href="#"
              className="bg-gradient-to-r from-orange-500 to-orange-800 py-2 px-3 rounded-md"
            >
              Create an account
            </a>
          </div>

          <div className="lg:hidden md:flex flex-col justify-end">
            <button onClick={toggleNavbar}>
              {mobileOpenDrawer ? <X /> : <Menu />}
            </button>
          </div>
        </div>

        {mobileOpenDrawer && (
          <div className="fixed z-50 right-0 flex-col justify-center items-center lg:hidden w-full bg-neutral-900 rounded-md opacity-95">
            <ul>
              {navItems.map((item, index) => (
                <li className="py-4 px-3 border-b decoration-slate-800 mx-3 my-1" key={index}>
                    {
                        item.label==="Social-Media" ? (
                            <div className="flex flex-shrink-0 items-center">
                            <img
                            className="h-6 w-6 bg-orange-600"
                            src={item.href}
                        />
                            <h4 className="ml-2 logo-font">Twitter</h4>
                            </div>
                        ):(
                            <a className="logo-font" href={item.href}>{item.label}</a>
                        )
                    }
                </li>
              ))}
            </ul>
          </div>
        )}
      </div>
    </nav>
  );
};

export default Navbar;
