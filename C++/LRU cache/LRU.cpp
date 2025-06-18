
#include<bits/stdc++.h>

struct Node{
    int key;
    Node* next;
    Node* prev;
    int val;

    Node(int key , int value){
        this->key = key;
        this->val = value;
        this->next = nullptr;
        this->prev = nullptr;
    };
};


class LRU{
    int capacity;
    std::unordered_map<int,Node*>mp;
    Node* start;
    Node* end;

    public:
    LRU(int cap);
    
    int get(int key);
    void put(int key , int value);
    void remove(int key);
};



LRU::LRU(int cap){
    capacity = cap;
    start->prev = nullptr;
    start->next = end;
    end->prev = start;
    end->next = nullptr;
}

int LRU::get(int key){
    if(mp.find(key)==mp.end()){
        return -1;
    }
    else{
        Node* node =mp[key];

        Node* prev = node->prev;
        Node* next = node->next;

        prev->next = node->next;
        next->prev = node->prev;

        node->next = start->next;
        node->prev = start;
        start->next = node;

        return node->val;
    }
}

void LRU::put(int key , int value){
    if(mp.find(key)!=mp.end()){
        get(key);
        start->next->val=value;
    }

    else {
        if(mp.size()==capacity){
        
            Node* lru = end->prev;
            get(lru->key);
            mp.erase(mp.find(key));
            mp[key] = lru;
        
        }
        else{
            Node* node = new Node(key , value);
            mp[key] = node;

            node->next = start->next;
            node->prev = start;
            start->next = node;
            node->next->prev = node; 
        
        }
    }
}
