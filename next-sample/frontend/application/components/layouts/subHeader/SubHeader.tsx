import Link from 'next/link';
import React from 'react';


const SubHeader = () => {

  return (
    <div className="sticky top-0 z-40 flex h-14 select-none items-center justify-center bg-[#f1eeef] shadow-md ">

        <div className="flex items-center justify-center px-6 sm:px-2">
          <Link href="/list?type=1">
            初期
          </Link>
        </div>
        <div className="flex items-center justify-center px-6 sm:px-2">
          <Link href="/list?type=2">
            中期
          </Link>
        </div>
        <div className="flex items-center justify-center px-6 sm:px-2">
          <Link href="/list?type=3">
            後期
          </Link>
        </div>
        <div className="flex items-center justify-center px-6 sm:px-2">
          <Link href="/list">
            全て
          </Link>
        </div>
      </div>
  );
};

export default SubHeader;