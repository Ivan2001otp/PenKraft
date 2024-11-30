// RecentPosts.js
import React from 'react';
import RecentPost from './RecentBlog'

const RecentPosts = ({ posts }) => {
  return (
    <div>
       <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-10 mt-8">
      {posts.map((post, index) => (
        <RecentPost key={index} {...post} />
      ))}
    </div>
    <div className='md:hidden block'>
    <a 
             onClick={()=>{console.log("View All")}}
            
             className="block md:hidden text-center mt-6 py-2 border-2 border-slate-100 mb-4 rounded-md">View All</a>
    </div>
    </div>
   
  );
};

export default RecentPosts;