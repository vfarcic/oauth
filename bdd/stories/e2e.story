


Meta:


Narrative:
In order to 
As a 
I want to 

GivenStories: 

Lifecycle:

Before:


After:



Scenario: Auth should redirect to Google service

Meta:


Given Web home page is opened
When Web user sets value john.doe@farcic.com to the element Email
When Web user sets value johndoe2 to the element Passwd
When Web user clicks the element signIn
When Web user clicks the element submit_approve_access
Then Web element fullName is present

Examples:
