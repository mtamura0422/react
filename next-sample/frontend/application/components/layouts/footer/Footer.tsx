

import { faHouse, faMagnifyingGlass, faPen } from "@fortawesome/free-solid-svg-icons";
import Link from 'next/link';
import { FooterParts } from './FooterParts'
import React from 'react';


const Footer = () => {

 

  return (
    
    <div className="sticky bottom-0 z-30 flex h-14 w-full select-none items-center justify-around -space-x-3 bg-[#f4cdd8] pt-0.5 text-center text-3xl">
      <Link href="/">
       <FooterParts pathname="/" text="ホーム" icon={faHouse} />
      </Link>
      <Link href="/searchrecipe">
        <FooterParts pathname="/searchrecipe" text="検索" icon={faMagnifyingGlass} />
      </Link>
  
      <Link href={'/regist'}>
        <FooterParts pathname="/regist" text="レシピ投稿" icon={faPen} />
      </Link>

      
    </div>
  );
};

export default Footer;