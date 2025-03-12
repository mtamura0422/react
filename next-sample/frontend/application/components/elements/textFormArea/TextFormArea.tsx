import { IconDefinition } from '@fortawesome/fontawesome-svg-core';

import React from 'react';
import Label from '@/components/elements/label/Label';
import TextArea from './TextArea';

type TextFormProps = {
  label?: string;
  labelIcon?: IconDefinition;
  placeholder?: string;
  value?: string | number;
  maxLength?: number;
  onChange?: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
};

const TextFormArea = React.memo((props: TextFormProps) => {
  const { placeholder } = props;

  return (
    <div className={`flex w-full mt-10 flex-col items-end`}>
      <div className="relative w-full min-w-[100px]">
        <TextArea
          placeholder={placeholder}
          value={props.value} // ここにvalueプロパティを追加（編集ページの初期値表示のために必要）
          onChange={props.onChange}
          maxLength={props.maxLength}
        />
        <Label 
          label={props.label}
          labelIcon={props.labelIcon}
        />
      </div>
    </div>
  );
});

export default TextFormArea;