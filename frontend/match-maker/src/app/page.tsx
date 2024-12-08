import { Match } from '@/types/types';

async function getMatches(): Promise<Match[]> {
  try {
    const response = await fetch('http://localhost:8080/matches', { 
      cache: 'no-store', // or 'no-cache' if you want to revalidate each request
      next: { revalidate: 10 } // Optional: revalidate every 10 seconds
    });

    if (!response.ok) {
      throw new Error('Failed to fetch matches');
    }

    return response.json();
  } catch (error) {
    console.error('Error fetching matches:', error);
    return [];
  }
}

export default async function Home() {
  const matches = await getMatches();

  return (
    <div className="min-h-screen bg-gray-100 py-10">
      <div className="container mx-auto px-4">
        <h1 className="text-4xl font-bold text-center mb-10 text-gray-800">
          Match Maker: Upcoming Football Matches
        </h1>

        {matches.length === 0 ? (
          <div className="text-center text-gray-600">
            <p className="text-2xl">No matches scheduled at the moment</p>
          </div>
        ) : (
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {matches.map((match) => (
              <div 
                key={match.Id} 
                className="bg-white rounded-lg shadow-md overflow-hidden transform transition duration-300 hover:scale-105"
              >
                <div className="p-6">
                  <div className="flex justify-between items-center mb-4">
                    <h2 className="text-2xl font-semibold text-gray-800">Match #{match.Id}</h2>
                    <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded">
                      Scheduled
                    </span>
                  </div>
                  
                  <div className="space-y-3 text-gray-600">
                    <div className="flex justify-between">
                      <span className="font-medium">Home Team:</span>
                      <span>Team ID {match.HomeTeamId}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="font-medium">Away Team:</span>
                      <span>Team ID {match.AwayTeamId}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="font-medium">Stadium:</span>
                      <span>Stadium ID {match.StadiumId}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="font-medium">Date:</span>
                      <span>{new Date(match.Date).toLocaleString()}</span>
                    </div>
                    <div className="border-t pt-3">
                      <h3 className="text-lg font-semibold mb-2 text-gray-700">Match Officials</h3>
                      <div className="flex justify-between">
                        <span className="font-medium">Main Referee:</span>
                        <span>{match.MainReferee}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="font-medium">Lineman 1:</span>
                        <span>{match.Lineman1}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="font-medium">Lineman 2:</span>
                        <span>{match.Lineman2}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}