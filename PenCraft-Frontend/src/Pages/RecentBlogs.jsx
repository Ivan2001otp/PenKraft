import React from 'react';
import RecentPost from './RecentBlog'
import { useNavigate } from 'react-router';

const RecentPosts = ({ posts }) => {
  const navigate = useNavigate();

  const handleReadMore=()=>{
    navigate('/read-more', {state:{list: posts}});
  };

  return (
    <div>
       <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-10 mt-8">
      {posts.map((post, index) => (
        <RecentPost key={index} {...post} />
      ))}
    </div>
    
    <div className='text-xl lg:text-2xl text-white flex  items-center justify-center'>
      <a 
      onClick={handleReadMore}
      className='bg-yellow-700 border-2 rounded-3xl px-3 py-2 transition-all duration-500 lg:hover:scale-125'>
        Read More
      </a>
    </div>
    </div>
   
  );
};

export default RecentPosts;