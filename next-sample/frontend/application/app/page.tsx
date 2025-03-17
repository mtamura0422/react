
import React from 'react';
import SubHeader from '@/components/layouts/subHeader/SubHeader';
import Pagination from '@/components/elements/pager/Pager';
// import List, {Recipe} from '@/components/elements/recipe/List';

import * as Recipe from "@/features/recipe/components/Index"

async function getData(page: number) {
 // var res = await fetch("http://localhost:9000/recipe/list")
 var res = await fetch("http://backend-service:9000/recipe/list?page=" + page)

 return res.json();
}

export default async function Home({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: number | undefined }>
}) {
  var { page = 1 } = await searchParams
 
  const data = await getData(page);
  const recipes: Recipe.RecipeData[] = data.list

  return (
<div>
  <SubHeader />
  <div className="mt-4">
    <div className="flex max-h-[400px] w-full flex-col overflow-y-scroll">
      {recipes.map((rp) => (
        <Recipe.List key={rp.id} id={rp.id} title={rp.title} content={rp.content} image={rp.image} materials={rp.materials} />
      ))}

      <Pagination
        page={page} 
        total_count={data.total_count} 
        per_page={data.per_page} 
        page_path=""
      />


    </div>
  </div>

</div>
  );
};
