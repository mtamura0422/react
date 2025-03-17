
import React from 'react';
import SubHeader from '@/components/layouts/subHeader/SubHeader';
// import List, {Recipe} from '@/components/elements/recipe/List';

import * as Recipe from "@/features/recipe/components/Index"

async function getData() {
 // var res = await fetch("http://localhost:9000/recipe/list")
 var res = await fetch("http://backend-service:9000/recipe/list")

 return res.json();
}



export default async function Home() {

  const recipes: Recipe.RecipeData[] = await getData();

  return (
<div>
  <SubHeader />
  <div className="mt-4">
    <div className="flex max-h-[400px] w-full flex-col overflow-y-scroll">
      {recipes.map((rp) => (
        <Recipe.List key={rp.id} id={rp.id} title={rp.title} content={rp.content} image={rp.image} materials={rp.materials} />
      ))}
    </div>
  </div>

</div>
  );
};
