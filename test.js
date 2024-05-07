'use client'

import React, { useEffect } from 'react';
import { autocomplete } from '@algolia/autocomplete-js';
import { createQuerySuggestionsPlugin } from '@algolia/autocomplete-plugin-query-suggestions';

import algoliasearch from 'algoliasearch/lite';
import '@algolia/autocomplete-theme-classic';

const appId = 'latency';
const apiKey = '6be0576ff61c053d5f9a3225e2a90f76';
const searchClient = algoliasearch(appId, apiKey);

const Search = () => {
  useEffect(() => {
    const querySuggestionsPlugin = createQuerySuggestionsPlugin({
      searchClient,
      indexName: 'instant_search_demo_query_suggestions',
      getSearchParams() {
        return {
          hitsPerPage: 10,
        };
      },
    });

    autocomplete({
      container: '#autocomplete',
      placeholder: 'Search',
      openOnFocus: true,
      plugins: [querySuggestionsPlugin],
    });
  }, []);  // Empty dependency array to ensure this runs only once after the component mounts

  return (
      <div className="container">
        <div id="autocomplete"></div>
      </div>
  );
};

export default Search;
