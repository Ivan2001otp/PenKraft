// RecentPost.js
import React from 'react';

const RecentPost = ({ title, date, description, image }) => {
  return (
    <div className="bg-neutral-700 shadow-md rounded-lg transition duration-300 ease-in-out hover:border-orange-500 hover:border-2 hover:bg-gradient-to-t hover:from-neutral-700 hover:via-emerald-500 hover:to-transparent-800 cursor-pointer mb-4">
      <img src={image} alt={title} className="w-full h-80 object-cover rounded-t-lg mb-4" />
      <h3 className="text-lg font-bold mb-2">{title}</h3>
      <p className="text-white-600 text-sm">{date}</p>
      <p className="text-gray-700">{description}</p>
    </div>
  );
};

export default RecentPost;