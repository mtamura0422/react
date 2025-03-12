import Link from 'next/link';
import React from 'react';


const Header = () => {
  return (
    <div className="sticky top-0 z-40 flex h-14 select-none items-center justify-center bg-[#f4cdd8] shadow-md">

        <div className="flex items-center justify-center">
          <Link href="/">
            離乳食レシピ
          </Link>
        </div>
      </div>
  );
};

export default Header;