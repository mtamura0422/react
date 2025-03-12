"use client"

import { usePathname } from 'next/navigation';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { IconDefinition } from '@fortawesome/fontawesome-svg-core';

export function FooterColor(props: {pathname: string}) {

  const pathname = usePathname();
  if (pathname === props.pathname) {
    return 'w-20 transition duration-75 ease-in-out hover:scale-105 ';
  }
  return ''

}


export function FooterParts(props: {pathname: string , text: string, icon: IconDefinition}) {

  const pathname = usePathname();
  // Footerアイコンなどの色を変える
  const GetFooterColor = () => {
    if (pathname === props.pathname) {
      return 'text-orange-400';
    }
    return ''
    
  };


  return (
    
    <div
      className={
        'w-20 transition duration-75 ease-in-out hover:scale-105 text-[#75665C] hover:text-orange-500 ' +
        GetFooterColor()
      }
    >
      <div>
        <FontAwesomeIcon icon={props.icon} className="m-auto h-8 w-8" />
        <div className="text-xs">
          {pathname === props.pathname ? props.text : ''}
        </div>
      </div>
    </div>
  )
  
}
