import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';

type LabelProps = {
  label?: React.ReactNode;
  labelIcon?: IconDefinition;
};

const Label = React.memo(({label = undefined, labelIcon = undefined}: LabelProps) => {
 
  if (label !== undefined || labelIcon != undefined) {
    
    return  (
      <label className="after:content['']
        peer-disabled:peer-placeholder-shown:text-blue-gray-500 
        pointer-events-none absolute -top-2.5 left-0 flex h-full w-full select-none text-base 
        font-semibold leading-tight transition-all 
        after:border-orange-500 after:transition-transform 
        after:duration-300 peer-placeholder-shown:leading-tight 
        peer-focus:text-sm peer-focus:leading-tight peer-focus:text-orange-500 
        peer-focus:after:scale-x-100 peer-focus:after:border-orange-500">
        <div>
          {labelIcon && (
            <FontAwesomeIcon icon={labelIcon} className="mr-1 w-5" />
          )}
          {label}
        </div>
      </label>
    )
  }
});

export default Label;