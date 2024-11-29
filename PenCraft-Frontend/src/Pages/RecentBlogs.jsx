// RecentPosts.js
import React from 'react';
import RecentPost from './RecentBlog'

const RecentPosts = ({ posts }) => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-10 mt-8">
      {posts.map((post, index) => (
        <RecentPost key={index} {...post} />
      ))}
    </div>
  );
};

export default RecentPosts;