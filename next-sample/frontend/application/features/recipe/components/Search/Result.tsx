
import * as Recipe from "@/features/recipe/components/Index"
import Pagination from '@/components/elements/pager/Pager';

type Props = {
  q: string;
  page: string;
};


const getData = async ({
  q,
  page,
}: Props) => {

  const params = { q: q, page: page };
  const query = new URLSearchParams(params);
    
  var res = await fetch("http://backend-service:9000/recipe/search?" + query)

  return res.json();

};

const Result = async ({ q, page }: Props) => {

  const data = await getData({ q, page });
  const recipes: Recipe.RecipeData[] = data.list
  
  
  return (
<div>
  <div className="mt-4">
    <div className="flex max-h-[400px] w-full flex-col overflow-y-scroll">
      {(recipes.length <= 0)? (
        <div className="text-center">該当するレシピが見つかりませんでした</div>
      ) : null}
      {recipes.map((rp) => (
        <Recipe.List key={rp.id} id={rp.id} title={rp.title} content={rp.content} image={rp.image} materials={rp.materials} />
      ))}
    </div>
    <Pagination
      page={Number(page)} 
      total_count={data.total_count} 
      per_page={data.per_page} 
      page_path="/search" 
      q={q}
    />
  </div>

</div>
  );
};
export default Result;