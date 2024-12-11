// RecentPost.js
import React from "react";

const RecentPost = ({ title, date, description, image }) => {
  return (
    <div className=" shadow-md rounded-lg transition duration-500 ease-in-out cursor-pointer mb-4">
      <img
        src={image}
        alt={title}
        className="w-full h-[280px] rounded-md border border-slate-400 object-cover rounded-t-lg mb-4"
      />
      <div className="-mt-12 mb-4  bg-green-600 absolute px-4 py-1 bg-opacity-80 rounded-r-xl z-10 ">
        {" "}
        Technology
      </div>

      <div className="p-2">
        <p className="text-sm text-slate-400 bokor-regular">{date}</p>
        <h3 className="text-lg font-bold mb-2 bokor-regular tracking-wider">{title}</h3>
        <p className="text-left text-slate-200 logo-font tracking-normal">{description}</p>
      </div>
    </div>
  );
};

export default RecentPost;
