# Todo List

## Add responses to database
This will allow for later translation and multiple language support

## Reimplement recipe search 
Currently, looking for a recipe returns the first recipe ingredients found.
Reimplement this so it returns recipes matching the input recipe name
Then, let the user reprompt for ingredients or step by step guide for one recipe

## Implement multiple ingredients input
Currently you can only prompt to search recipes that use a single ingredient.
Implement logic to receive a list of ingredients and search for recipes using all of them

## Reprompt for random recipe
Allow the user to remprompt for ingredients or guide when a random recipe is returned

## Database overhaul
Normalize ingredients structure, make it a set of maps {"ingredient": "name", "amount": "value"} so matching is easier
Normalize step by step to allow prompting for one step at a time. Make it an ordered list.
Make the status persists, so the user can interact with Alexa while cooking, but at the same time make him/her able to ask for new directions

## Multilingual support
This, multilingual support, maybe let the database as it is and translate with an LLM? That would be fun to implement <3

## Integrate with fridge list
If the user has a fridge list with all its assets, integrate with it so the skills is able to say which ingredients are missing
Or to allow searching for recipes which the user has all ingredients at hand
