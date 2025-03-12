import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import React from 'react';

import Input from './Input';
import Label from '@/components/elements/label/Label';

type TextFormProps = {
  label?: React.ReactNode;
  labelIcon?: IconDefinition;
  placeholder?: string;
  value?: string | number;
  type?: string;
  min?: number;
  max?: number;
  maxLength?: number;
  required?: boolean;
  disabled?: boolean;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const TextForm = React.memo((props: TextFormProps) => {
  const { placeholder, min, max } = props;

  return (
    <div className={`flex w-full mx-auto mt-10`}>
      <div className="h-15 relative w-full">
        <Input
          placeholder={placeholder}
          value={props.value} // ここにvalueプロパティを追加（編集ページの初期値表示のために必要）
          type={props.type}
          min={props.type === 'number' ? min : undefined}
          max={props.type === 'number' ? max : undefined}
          onChange={props.onChange}
          maxLength={props.maxLength}
          required={props.required}
          disabled={props.disabled}
        />
  
        <Label 
          label={props.label}
          labelIcon={props.labelIcon}
        />
        
      </div>
    </div>
  );
});

export default TextForm;