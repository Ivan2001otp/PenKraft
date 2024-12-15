// RecentPost.js
import React from "react";

const formatDateString=(dateString)=>{
  const date = new Date(dateString)
  const getDayWithSuffix=(day)=>{
    const suffix = ['th', 'st', 'nd', 'rd'];
    const val = day % 100;
    return day + (suffix[(val-20) % 10] || suffix[val] || suffix[0])

  };

  const day = getDayWithSuffix(date.getDate())
  const month = date.toLocaleString('default', { month: 'short' }); // Get the abbreviated month (e.g., "Dec")
  const year = date.getFullYear(); 

  return `${day} ${month}, ${year}`;
}

const RecentPost = ({ blog_id,body ,excerpt, image, slug, tag_name, tag_id, updated_at ,user_id ,title  }) => {

 
  return (
    <div className=" shadow-md rounded-lg transition duration-500 ease-in-out cursor-pointer mb-4 hover:border-2 hover:border-slate-300 hover:shadow-xl hover:translate-y-[-4px] lg:hover:scale-105">
      <img
        src={image}
        alt={title}
        className="w-full h-[280px] rounded-md border border-slate-400 object-cover rounded-t-lg mb-4"
      />
      <div className="-mt-12 mb-4  bg-green-600 absolute px-4 py-1 bg-opacity-80 rounded-r-xl z-10 ">
        {" "}
        {tag_name}
      </div>

      <div className="p-2">
        <p className="text-sm text-slate-400 logo-font mb-2">{formatDateString(updated_at)}</p>
        <h3 className="text-lg font-bold mb-2 bokor-regular tracking-wider">{title}</h3>
        <p className="text-left text-slate-200 logo-font tracking-normal">{excerpt}</p>
      </div>
    </div>
  );
};

export default RecentPost;
