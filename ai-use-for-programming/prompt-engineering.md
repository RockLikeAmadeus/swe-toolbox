# Tips from AI-Powered Developer

- "the _Refinement Pattern in prompt engineering_: iteratively refining or improving the prompt to get more accurate, relevant, or sophisticated responses."
	- This pattern can be better implemented by feeding the chat something like "From now on, when I give you a prompt, output a better prompt."
- "the Persona Pattern refers to a strategy of designing prompts that establish a specific persona or role for the AI to assume."
	- "You can apply the Persona Pattern in either direction: you can tell the LLM to respond as though it were someone or something within a given role, or you can ask the LLM to assume that you are a certain persona."
		- "When you use the Persona Pattern in reverse, it is commonly referred to as the _Audience Persona Pattern_ in the context of prompt engineering."



# Blurbs for Prompt COnstruction, from AI-Powered Developer

- After requesting a high-level design of a software project, or even simply describing the purpose and requirements of the application: "Show the Python package structure for this application."
	- "Please show the source code for it_asset_management/app/schemas/asset.py."
	- "I would like to build an ITAM project written in Python. It will focus on the tracking and management of Hardware. It should expose REST APIs, using FastAPI, and persist data using SQLAlchemy. It should use hexagonal architecture. As a software architect, please show me the Mermaid class diagram for this project.""