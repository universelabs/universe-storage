# go-backend

Universe back end in Go.  All documentation here is first-draft and planning oriented.  It will likely change.  

# First step: API server
Http server that works quite like the one that Dan Trevino built in Express-- in fact hopefully exactly, with the exception of user auth, for now.  

# Second Step:  Generate mock keypars
(question: Should we be doing this on the front end, anyway-- feeding them to the back end?).  Need to look into libraries for generating Blockstack, Bitcoin, and Ethereum addresses (on either the front or the back in fact-- main thing is that we want 

# Third Step: Serve the front end from the back end

When the user runs this, it shouldn't be just a plain REST API, that wouldn't give them too much to chew on at all.  Instead, we should embed the front end into go-backend so that when this binary is run, the user can go to a certain port and have the front end served to them.  Default should be localhost (mac, windows) and there should be a flag to serve at 0.0.0.0 (Raspberry pi).  


## Fourth Step: Persistence
* Go Embedded DB
  * Boltdb
  * bbolt
  * storm
  * badgerdb
  
Rationale: This way, it runs anywhere and creates a new database at first launch, while using a fixed location so that subsequent launches pick up anything persisted before.
