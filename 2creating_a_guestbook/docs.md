https://cloud.google.com/appengine/docs/go/gettingstarted/usingdatastore


The Datastore is extremely resilient in the face of catastrophic failure, but its consistency guarantees may differ from what you're familiar with.


Entities descended from a common ancestor
are said to belong to the same entity group;

the common ancestor's key is the group's parent key,
which serves to identify the entire group.



Queries over a single entity group,
called ancestor queries, refer to the parent key instead of a specific entity's key.



Entity groups are a unit of both consistency and transactionality:


whereas queries over multiple entity groups
may return stale, eventually consistent results, those limited to a single entity group always return up-to-date, strongly consistent results.

--------------------------------------------------------------------------------------------------------------------------------------------------
Because querying in the High Replication Datastore is strongly consistent only within entity groups,
we assign all of one book's greetings to the same entity group in this example
by setting the same parent for each greeting.


This means a user will always see a greeting immediately after it was written.


However, the rate at which you can write to the same entity group is limited to 1 write to the entity group per second.


When you design a real application you'll need to keep this fact in mind.


By using services such as Memcache,
you can mitigate the chance that a user won't see fresh results
when querying across entity groups immediately after a write.
