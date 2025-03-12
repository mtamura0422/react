import Link from 'next/link';
import React, { useEffect } from 'react';


export type RecipeData = {
  id: number;
  title: string;
  content: string;
  image: string;
  materials: [];
  updated_at: string;
};


export default async function RecipeList(props: RecipeData) {
  
  var { id, title, content, image, materials } = props;

  if (image == undefined) {
    image = "/images/noimage.png"
  }

  return (
      <div>
          <div
            key={id}
            className="mb-3 h-28 rounded-lg bg-[#ffffff] py-2 shadow-md hover:bg-[#f7e5f0] hover:text-pink-500 hover:shadow-lg"
          >
            <Link href={'/detail/' + id}>
              <div>
                <div className="flex">
                  <div className="flex-shrink-0">
                    <img
                      className="mx-1.5 h-24 w-32 rounded-lg border-2 border-solid border-[#f0bee2] object-cover lg:mx-2"
                      alt="レシピリスト"
                      src={image}
                    />
                  </div>
                  <div className="flex w-full flex-col">
                    <div className="mr-1 line-clamp-1 break-all font-bold">
                      {title}
                    </div>
                    <div className="my-auto ml-0.5 mr-1 line-clamp-2 overflow-hidden whitespace-pre-wrap break-all text-xs">
                      {content}
                    </div>
                    <div className="mr-1 text-left text-xs font-semibold lg:mr-2 lg:text-base">
                    <span className="mr-1">材料 :</span>
                    {materials.map((material) =>
                      <span className="mr-1">{material}</span>
                    )}

                    </div>
                  </div>
                </div>
              </div>
            </Link>
          </div>
    </div>
  );
};
